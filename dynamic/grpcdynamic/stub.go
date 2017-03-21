package grpcdynamic

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

// Stub is an RPC client stub, used for dynamically dispatching RPCs to a server.
type Stub struct {
	conn *grpc.ClientConn
	er   *dynamic.ExtensionRegistry
}

// NewStub creates a new RPC stub that uses the given connection for communicating with a server.
func NewStub(conn *grpc.ClientConn) *Stub {
	return NewStubWithExtensionRegistry(conn, nil)
}

// NewStub creates a new RPC stub that uses the given connection for communicating with a server and the
// given ExtensionRegistry when parsing responses.
func NewStubWithExtensionRegistry(conn *grpc.ClientConn, er *dynamic.ExtensionRegistry) *Stub {
	return &Stub{conn, er}
}

// InvokeRpc sends a unary RPC and returns the response. Use this for unary methods.
func (s Stub) InvokeRpc(ctx context.Context, method *desc.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (*dynamic.Message, error) {
	if method.IsClientStreaming() || method.IsServerStreaming() {
		return fmt.Errorf("InvokeRpc is for unary methods; %q is %s", method.GetFullyQualifiedName(), methodType(method))
	}
	if err := checkMessageType(method.GetInputType(), request); err != nil {
		return err
	}
	resp := dynamic.NewMessageWithExtensionRegistry(method.GetOutputType(), s.er)
	if err := grpc.Invoke(ctx, method.GetFullyQualifiedName(), request, resp, s.conn, opts); err != nil {
		return nil, err
	}
	return resp, nil
}

// InvokeRpcStream sends a unary RPC and returns the response stream. Use this for server-streaming methods.
func (s Stub) InvokeRpcStream(ctx context.Context, method *desc.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (*Stream, error) {
	if method.IsClientStreaming() || !method.IsServerStreaming() {
		return fmt.Errorf("InvokeRpcStream is for server-streaming methods; %q is %s", method.GetFullyQualifiedName(), methodType(method))
	}
	if err := checkMessageType(method.GetInputType(), request); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	sd := grpc.StreamDesc{
		StreamName: method.GetFullyQualifiedName(),
		ServerStreams: method.IsServerStreaming(),
		ClientStreams: method.IsClientStreaming(),
	}
	if cs, err := grpc.NewClientStream(ctx, &sd, s.conn, method.GetFullyQualifiedName(), opts); err != nil {
		return nil, err
	} else {
		err = cs.SendMsg(request)
		if err != nil {
			cancel()
			return nil, err
		}
		err = cs.CloseSend()
		if err != nil {
			cancel()
			return nil, err
		}
		return &Stream{cs, method.GetOutputType(), s.er}, nil
	}
}

// NewClientStream creates a new stream that is used to both send request messages and receive response
// messages. Use this for methods that accept a stream of requests (e.g. client-streaming and bidi-streaming
// methods).
func (s Stub) NewClientStream(ctx context.Context, method *desc.MethodDescriptor, opts ...grpc.CallOption) (*ClientStream, error) {
	if !method.IsClientStreaming() {
		return fmt.Errorf("InvokeRpcStream is for client- and bidi-streaming methods; %q is %s", method.GetFullyQualifiedName(), methodType(method))
	}
	sd := grpc.StreamDesc{
		StreamName: method.GetFullyQualifiedName(),
		ServerStreams: method.IsServerStreaming(),
		ClientStreams: method.IsClientStreaming(),
	}
	if cs, err := grpc.NewClientStream(ctx, &sd, s.conn, method.GetFullyQualifiedName(), opts); err != nil {
		return nil, err
	} else {
		return &ClientStream{Stream{cs, method.GetOutputType(), s.er}, method.GetInputType()}, nil
	}
}

func methodType(md *desc.MethodDescriptor) string {
	if md.IsClientStreaming() && md.IsClientStreaming() {
		return "bidi-streaming"
	} else if md.IsClientStreaming() {
		return "client-streaming"
	} else if md.IsServerStreaming() {
		return "server-streaming"
	} else {
		return "unary"
	}
}

func checkMessageType(md *desc.MessageDescriptor, msg proto.Message) error {
	var typeName string
	if dm, ok := msg.(*dynamic.Message); ok {
		typeName = dm.GetMessageDescriptor().GetFullyQualifiedName()
	} else {
		typeName = proto.MessageName(msg)
	}
	if typeName != md.GetFullyQualifiedName() {
		return fmt.Errorf("Expecting message of type %s; got %s", md.GetFullyQualifiedName(), typeName)
	}
	return nil
}

// Stream represents a response stream from a server. Messages in the stream can be queried as can
// header and trailer metadata sent by the server.
type Stream struct {
	stream   grpc.ClientStream
	respType *desc.MessageDescriptor
	er       *dynamic.ExtensionRegistry
}

// Header returns any header metadata sent by the server (blocks if necessary until headers are
// received).
func (s *Stream) Header() (metadata.MD, error) {
	return s.stream.Header()
}

// Trailer returns the trailer metadata sent by the server. It must only be called after
// RecvMsg returns a non-nil error (which may be EOF for normal completion of stream).
func (s *Stream) Trailer() (metadata.MD) {
	return s.stream.Trailer()
}

// Context returns the context associated with this streaming operation.
func (s *Stream) Context() context.Context {
	return
}

// RecvMsg returns the next message in the response stream or an error. If the stream
// has completed normally, the error is io.EOF. Otherwise, the error indicates the
// nature of the abnormal termination of the stream.
func (s *Stream) RecvMsg() (*dynamic.Message, error) {
	resp := dynamic.NewMessageWithExtensionRegistry(s.respType, s.er)
	if err := s.stream.RecvMsg(resp); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

// ClientStream represents a bi-directional stream for sending messages to and receiving
// messages from a server. For client-streaming operations (e.g. unary response), the
// server will send back only one message, after which subsequent invocations of RecvMsg
// return io.EOF.
type ClientStream struct {
	Stream
	reqType *desc.MessageDescriptor
}

// CloseSend indicates the request stream has ended. Invoke this after all request messages
// are sent (even if there are zero such messages).
func (cs *ClientStream) CloseSend() error {
	return cs.stream.CloseSend()
}

// SendMsg sends a request message to the server.
func (cs *ClientStream) SendMsg(m proto.Message) error {
	if err := checkMessageType(cs.reqType, m); err != nil {
		return err
	}
	return cs.stream.SendMsg(m)
}

