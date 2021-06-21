// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package prcontent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	GetSharedMedia(ctx context.Context, in *SharedMediaRequest, opts ...grpc.CallOption) (Content_GetSharedMediaClient, error)
	//rpc GetPostsByUser(GetPostsRequest) returns(GetPostsResponse);
	GetPostsByUser(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (Content_GetPostsByUserClient, error)
	GetStoriesIndividual(ctx context.Context, in *GetStoriesIndividualRequest, opts ...grpc.CallOption) (*GetStoriesIndividualResponse, error)
	GetProfilePicture(ctx context.Context, in *GetProfilePictureRequest, opts ...grpc.CallOption) (*GetProfilePictureResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (Content_GetCommentsClient, error)
	GetReactions(ctx context.Context, in *GetReactionsRequest, opts ...grpc.CallOption) (Content_GetReactionsClient, error)
	GetPostsByUserReaction(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (Content_GetPostsByUserReactionClient, error)
	GetHighlights(ctx context.Context, in *GetHighlightsRequest, opts ...grpc.CallOption) (*GetHighlightsResponse, error)
	CreateStory(ctx context.Context, in *CreateStoryRequest, opts ...grpc.CallOption) (*CreateStoryResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	AddPost(ctx context.Context, opts ...grpc.CallOption) (Content_AddPostClient, error)
	AddStory(ctx context.Context, opts ...grpc.CallOption) (Content_AddStoryClient, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
	AddReaction(ctx context.Context, in *AddReactionRequest, opts ...grpc.CallOption) (*AddReactionResponse, error)
	AddProfilePicture(ctx context.Context, opts ...grpc.CallOption) (Content_AddProfilePictureClient, error)
	AddHighlight(ctx context.Context, in *AddHighlightRequest, opts ...grpc.CallOption) (*AddHighlightResponse, error)
	PutReaction(ctx context.Context, in *PutReactionRequest, opts ...grpc.CallOption) (*PutReactionResponse, error)
	CreateUserFolder(ctx context.Context, in *CreateUserFolderRequest, opts ...grpc.CallOption) (*CreateUserFolderResponse, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) GetSharedMedia(ctx context.Context, in *SharedMediaRequest, opts ...grpc.CallOption) (Content_GetSharedMediaClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[0], "/Content/GetSharedMedia", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentGetSharedMediaClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Content_GetSharedMediaClient interface {
	Recv() (*SharedMediaResponse, error)
	grpc.ClientStream
}

type contentGetSharedMediaClient struct {
	grpc.ClientStream
}

func (x *contentGetSharedMediaClient) Recv() (*SharedMediaResponse, error) {
	m := new(SharedMediaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) GetPostsByUser(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (Content_GetPostsByUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[1], "/Content/GetPostsByUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentGetPostsByUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Content_GetPostsByUserClient interface {
	Recv() (*GetPostsResponse, error)
	grpc.ClientStream
}

type contentGetPostsByUserClient struct {
	grpc.ClientStream
}

func (x *contentGetPostsByUserClient) Recv() (*GetPostsResponse, error) {
	m := new(GetPostsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) GetStoriesIndividual(ctx context.Context, in *GetStoriesIndividualRequest, opts ...grpc.CallOption) (*GetStoriesIndividualResponse, error) {
	out := new(GetStoriesIndividualResponse)
	err := c.cc.Invoke(ctx, "/Content/GetStoriesIndividual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetProfilePicture(ctx context.Context, in *GetProfilePictureRequest, opts ...grpc.CallOption) (*GetProfilePictureResponse, error) {
	out := new(GetProfilePictureResponse)
	err := c.cc.Invoke(ctx, "/Content/GetProfilePicture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (Content_GetCommentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[2], "/Content/GetComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentGetCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Content_GetCommentsClient interface {
	Recv() (*GetCommentsResponse, error)
	grpc.ClientStream
}

type contentGetCommentsClient struct {
	grpc.ClientStream
}

func (x *contentGetCommentsClient) Recv() (*GetCommentsResponse, error) {
	m := new(GetCommentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) GetReactions(ctx context.Context, in *GetReactionsRequest, opts ...grpc.CallOption) (Content_GetReactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[3], "/Content/GetReactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentGetReactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Content_GetReactionsClient interface {
	Recv() (*GetReactionsResponse, error)
	grpc.ClientStream
}

type contentGetReactionsClient struct {
	grpc.ClientStream
}

func (x *contentGetReactionsClient) Recv() (*GetReactionsResponse, error) {
	m := new(GetReactionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) GetPostsByUserReaction(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (Content_GetPostsByUserReactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[4], "/Content/GetPostsByUserReaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentGetPostsByUserReactionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Content_GetPostsByUserReactionClient interface {
	Recv() (*GetPostsResponse, error)
	grpc.ClientStream
}

type contentGetPostsByUserReactionClient struct {
	grpc.ClientStream
}

func (x *contentGetPostsByUserReactionClient) Recv() (*GetPostsResponse, error) {
	m := new(GetPostsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) GetHighlights(ctx context.Context, in *GetHighlightsRequest, opts ...grpc.CallOption) (*GetHighlightsResponse, error) {
	out := new(GetHighlightsResponse)
	err := c.cc.Invoke(ctx, "/Content/GetHighlights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateStory(ctx context.Context, in *CreateStoryRequest, opts ...grpc.CallOption) (*CreateStoryResponse, error) {
	out := new(CreateStoryResponse)
	err := c.cc.Invoke(ctx, "/Content/CreateStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/Content/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddPost(ctx context.Context, opts ...grpc.CallOption) (Content_AddPostClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[5], "/Content/AddPost", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentAddPostClient{stream}
	return x, nil
}

type Content_AddPostClient interface {
	Send(*AddPostRequest) error
	CloseAndRecv() (*AddPostResponse, error)
	grpc.ClientStream
}

type contentAddPostClient struct {
	grpc.ClientStream
}

func (x *contentAddPostClient) Send(m *AddPostRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentAddPostClient) CloseAndRecv() (*AddPostResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddPostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) AddStory(ctx context.Context, opts ...grpc.CallOption) (Content_AddStoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[6], "/Content/AddStory", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentAddStoryClient{stream}
	return x, nil
}

type Content_AddStoryClient interface {
	Send(*AddStoryRequest) error
	CloseAndRecv() (*AddStoryResponse, error)
	grpc.ClientStream
}

type contentAddStoryClient struct {
	grpc.ClientStream
}

func (x *contentAddStoryClient) Send(m *AddStoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentAddStoryClient) CloseAndRecv() (*AddStoryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddStoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	out := new(AddCommentResponse)
	err := c.cc.Invoke(ctx, "/Content/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddReaction(ctx context.Context, in *AddReactionRequest, opts ...grpc.CallOption) (*AddReactionResponse, error) {
	out := new(AddReactionResponse)
	err := c.cc.Invoke(ctx, "/Content/AddReaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddProfilePicture(ctx context.Context, opts ...grpc.CallOption) (Content_AddProfilePictureClient, error) {
	stream, err := c.cc.NewStream(ctx, &Content_ServiceDesc.Streams[7], "/Content/AddProfilePicture", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentAddProfilePictureClient{stream}
	return x, nil
}

type Content_AddProfilePictureClient interface {
	Send(*AddProfilePictureRequest) error
	CloseAndRecv() (*AddProfilePictureResponse, error)
	grpc.ClientStream
}

type contentAddProfilePictureClient struct {
	grpc.ClientStream
}

func (x *contentAddProfilePictureClient) Send(m *AddProfilePictureRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentAddProfilePictureClient) CloseAndRecv() (*AddProfilePictureResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddProfilePictureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentClient) AddHighlight(ctx context.Context, in *AddHighlightRequest, opts ...grpc.CallOption) (*AddHighlightResponse, error) {
	out := new(AddHighlightResponse)
	err := c.cc.Invoke(ctx, "/Content/AddHighlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) PutReaction(ctx context.Context, in *PutReactionRequest, opts ...grpc.CallOption) (*PutReactionResponse, error) {
	out := new(PutReactionResponse)
	err := c.cc.Invoke(ctx, "/Content/PutReaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateUserFolder(ctx context.Context, in *CreateUserFolderRequest, opts ...grpc.CallOption) (*CreateUserFolderResponse, error) {
	out := new(CreateUserFolderResponse)
	err := c.cc.Invoke(ctx, "/Content/CreateUserFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	GetSharedMedia(*SharedMediaRequest, Content_GetSharedMediaServer) error
	//rpc GetPostsByUser(GetPostsRequest) returns(GetPostsResponse);
	GetPostsByUser(*GetPostsRequest, Content_GetPostsByUserServer) error
	GetStoriesIndividual(context.Context, *GetStoriesIndividualRequest) (*GetStoriesIndividualResponse, error)
	GetProfilePicture(context.Context, *GetProfilePictureRequest) (*GetProfilePictureResponse, error)
	GetComments(*GetCommentsRequest, Content_GetCommentsServer) error
	GetReactions(*GetReactionsRequest, Content_GetReactionsServer) error
	GetPostsByUserReaction(*GetPostsRequest, Content_GetPostsByUserReactionServer) error
	GetHighlights(context.Context, *GetHighlightsRequest) (*GetHighlightsResponse, error)
	CreateStory(context.Context, *CreateStoryRequest) (*CreateStoryResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	AddPost(Content_AddPostServer) error
	AddStory(Content_AddStoryServer) error
	AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error)
	AddReaction(context.Context, *AddReactionRequest) (*AddReactionResponse, error)
	AddProfilePicture(Content_AddProfilePictureServer) error
	AddHighlight(context.Context, *AddHighlightRequest) (*AddHighlightResponse, error)
	PutReaction(context.Context, *PutReactionRequest) (*PutReactionResponse, error)
	CreateUserFolder(context.Context, *CreateUserFolderRequest) (*CreateUserFolderResponse, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) GetSharedMedia(*SharedMediaRequest, Content_GetSharedMediaServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSharedMedia not implemented")
}
func (UnimplementedContentServer) GetPostsByUser(*GetPostsRequest, Content_GetPostsByUserServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPostsByUser not implemented")
}
func (UnimplementedContentServer) GetStoriesIndividual(context.Context, *GetStoriesIndividualRequest) (*GetStoriesIndividualResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoriesIndividual not implemented")
}
func (UnimplementedContentServer) GetProfilePicture(context.Context, *GetProfilePictureRequest) (*GetProfilePictureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfilePicture not implemented")
}
func (UnimplementedContentServer) GetComments(*GetCommentsRequest, Content_GetCommentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedContentServer) GetReactions(*GetReactionsRequest, Content_GetReactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetReactions not implemented")
}
func (UnimplementedContentServer) GetPostsByUserReaction(*GetPostsRequest, Content_GetPostsByUserReactionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPostsByUserReaction not implemented")
}
func (UnimplementedContentServer) GetHighlights(context.Context, *GetHighlightsRequest) (*GetHighlightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHighlights not implemented")
}
func (UnimplementedContentServer) CreateStory(context.Context, *CreateStoryRequest) (*CreateStoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStory not implemented")
}
func (UnimplementedContentServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedContentServer) AddPost(Content_AddPostServer) error {
	return status.Errorf(codes.Unimplemented, "method AddPost not implemented")
}
func (UnimplementedContentServer) AddStory(Content_AddStoryServer) error {
	return status.Errorf(codes.Unimplemented, "method AddStory not implemented")
}
func (UnimplementedContentServer) AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedContentServer) AddReaction(context.Context, *AddReactionRequest) (*AddReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReaction not implemented")
}
func (UnimplementedContentServer) AddProfilePicture(Content_AddProfilePictureServer) error {
	return status.Errorf(codes.Unimplemented, "method AddProfilePicture not implemented")
}
func (UnimplementedContentServer) AddHighlight(context.Context, *AddHighlightRequest) (*AddHighlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHighlight not implemented")
}
func (UnimplementedContentServer) PutReaction(context.Context, *PutReactionRequest) (*PutReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutReaction not implemented")
}
func (UnimplementedContentServer) CreateUserFolder(context.Context, *CreateUserFolderRequest) (*CreateUserFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserFolder not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_GetSharedMedia_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SharedMediaRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContentServer).GetSharedMedia(m, &contentGetSharedMediaServer{stream})
}

type Content_GetSharedMediaServer interface {
	Send(*SharedMediaResponse) error
	grpc.ServerStream
}

type contentGetSharedMediaServer struct {
	grpc.ServerStream
}

func (x *contentGetSharedMediaServer) Send(m *SharedMediaResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Content_GetPostsByUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContentServer).GetPostsByUser(m, &contentGetPostsByUserServer{stream})
}

type Content_GetPostsByUserServer interface {
	Send(*GetPostsResponse) error
	grpc.ServerStream
}

type contentGetPostsByUserServer struct {
	grpc.ServerStream
}

func (x *contentGetPostsByUserServer) Send(m *GetPostsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Content_GetStoriesIndividual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoriesIndividualRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetStoriesIndividual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/GetStoriesIndividual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetStoriesIndividual(ctx, req.(*GetStoriesIndividualRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetProfilePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfilePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetProfilePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/GetProfilePicture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetProfilePicture(ctx, req.(*GetProfilePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContentServer).GetComments(m, &contentGetCommentsServer{stream})
}

type Content_GetCommentsServer interface {
	Send(*GetCommentsResponse) error
	grpc.ServerStream
}

type contentGetCommentsServer struct {
	grpc.ServerStream
}

func (x *contentGetCommentsServer) Send(m *GetCommentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Content_GetReactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetReactionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContentServer).GetReactions(m, &contentGetReactionsServer{stream})
}

type Content_GetReactionsServer interface {
	Send(*GetReactionsResponse) error
	grpc.ServerStream
}

type contentGetReactionsServer struct {
	grpc.ServerStream
}

func (x *contentGetReactionsServer) Send(m *GetReactionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Content_GetPostsByUserReaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPostsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContentServer).GetPostsByUserReaction(m, &contentGetPostsByUserReactionServer{stream})
}

type Content_GetPostsByUserReactionServer interface {
	Send(*GetPostsResponse) error
	grpc.ServerStream
}

type contentGetPostsByUserReactionServer struct {
	grpc.ServerStream
}

func (x *contentGetPostsByUserReactionServer) Send(m *GetPostsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Content_GetHighlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHighlightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetHighlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/GetHighlights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetHighlights(ctx, req.(*GetHighlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/CreateStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateStory(ctx, req.(*CreateStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddPost_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentServer).AddPost(&contentAddPostServer{stream})
}

type Content_AddPostServer interface {
	SendAndClose(*AddPostResponse) error
	Recv() (*AddPostRequest, error)
	grpc.ServerStream
}

type contentAddPostServer struct {
	grpc.ServerStream
}

func (x *contentAddPostServer) SendAndClose(m *AddPostResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentAddPostServer) Recv() (*AddPostRequest, error) {
	m := new(AddPostRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Content_AddStory_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentServer).AddStory(&contentAddStoryServer{stream})
}

type Content_AddStoryServer interface {
	SendAndClose(*AddStoryResponse) error
	Recv() (*AddStoryRequest, error)
	grpc.ServerStream
}

type contentAddStoryServer struct {
	grpc.ServerStream
}

func (x *contentAddStoryServer) SendAndClose(m *AddStoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentAddStoryServer) Recv() (*AddStoryRequest, error) {
	m := new(AddStoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Content_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/AddReaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddReaction(ctx, req.(*AddReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddProfilePicture_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentServer).AddProfilePicture(&contentAddProfilePictureServer{stream})
}

type Content_AddProfilePictureServer interface {
	SendAndClose(*AddProfilePictureResponse) error
	Recv() (*AddProfilePictureRequest, error)
	grpc.ServerStream
}

type contentAddProfilePictureServer struct {
	grpc.ServerStream
}

func (x *contentAddProfilePictureServer) SendAndClose(m *AddProfilePictureResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentAddProfilePictureServer) Recv() (*AddProfilePictureRequest, error) {
	m := new(AddProfilePictureRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Content_AddHighlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHighlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddHighlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/AddHighlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddHighlight(ctx, req.(*AddHighlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_PutReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).PutReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/PutReaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).PutReaction(ctx, req.(*PutReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateUserFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateUserFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Content/CreateUserFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateUserFolder(ctx, req.(*CreateUserFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStoriesIndividual",
			Handler:    _Content_GetStoriesIndividual_Handler,
		},
		{
			MethodName: "GetProfilePicture",
			Handler:    _Content_GetProfilePicture_Handler,
		},
		{
			MethodName: "GetHighlights",
			Handler:    _Content_GetHighlights_Handler,
		},
		{
			MethodName: "CreateStory",
			Handler:    _Content_CreateStory_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Content_CreatePost_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Content_AddComment_Handler,
		},
		{
			MethodName: "AddReaction",
			Handler:    _Content_AddReaction_Handler,
		},
		{
			MethodName: "AddHighlight",
			Handler:    _Content_AddHighlight_Handler,
		},
		{
			MethodName: "PutReaction",
			Handler:    _Content_PutReaction_Handler,
		},
		{
			MethodName: "CreateUserFolder",
			Handler:    _Content_CreateUserFolder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSharedMedia",
			Handler:       _Content_GetSharedMedia_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPostsByUser",
			Handler:       _Content_GetPostsByUser_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetComments",
			Handler:       _Content_GetComments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetReactions",
			Handler:       _Content_GetReactions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPostsByUserReaction",
			Handler:       _Content_GetPostsByUserReaction_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddPost",
			Handler:       _Content_AddPost_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddStory",
			Handler:       _Content_AddStory_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddProfilePicture",
			Handler:       _Content_AddProfilePicture_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "content/content.proto",
}
