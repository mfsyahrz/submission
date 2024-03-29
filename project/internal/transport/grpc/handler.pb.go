// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: handler.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{0}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type MoviePagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *MovieMeta `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MoviePagination) Reset() {
	*x = MoviePagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviePagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviePagination) ProtoMessage() {}

func (x *MoviePagination) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviePagination.ProtoReflect.Descriptor instead.
func (*MoviePagination) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{2}
}

func (x *MoviePagination) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MoviePagination) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MoviePagination) GetData() *MovieMeta {
	if x != nil {
		return x.Data
	}
	return nil
}

type MovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *MovieDetail `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MovieResponse) Reset() {
	*x = MovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResponse) ProtoMessage() {}

func (x *MovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResponse.ProtoReflect.Descriptor instead.
func (*MovieResponse) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{3}
}

func (x *MovieResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MovieResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MovieResponse) GetData() *MovieDetail {
	if x != nil {
		return x.Data
	}
	return nil
}

type MovieMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	TotalResult string   `protobuf:"bytes,2,opt,name=totalResult,proto3" json:"totalResult,omitempty"`
	List        []*Movie `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MovieMeta) Reset() {
	*x = MovieMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieMeta) ProtoMessage() {}

func (x *MovieMeta) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieMeta.ProtoReflect.Descriptor instead.
func (*MovieMeta) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{4}
}

func (x *MovieMeta) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *MovieMeta) GetTotalResult() string {
	if x != nil {
		return x.TotalResult
	}
	return ""
}

func (x *MovieMeta) GetList() []*Movie {
	if x != nil {
		return x.List
	}
	return nil
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbId string `protobuf:"bytes,3,opt,name=imdbId,proto3" json:"imdbId,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{5}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Movie) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *Movie) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Movie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type MovieDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year       string         `protobuf:"bytes,1,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string         `protobuf:"bytes,2,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string         `protobuf:"bytes,3,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string         `protobuf:"bytes,4,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string         `protobuf:"bytes,5,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string         `protobuf:"bytes,6,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string         `protobuf:"bytes,7,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string         `protobuf:"bytes,8,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string         `protobuf:"bytes,9,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string         `protobuf:"bytes,10,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string         `protobuf:"bytes,11,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string         `protobuf:"bytes,12,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string         `protobuf:"bytes,13,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Metascore  string         `protobuf:"bytes,14,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating string         `protobuf:"bytes,15,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string         `protobuf:"bytes,16,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string         `protobuf:"bytes,17,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string         `protobuf:"bytes,18,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string         `protobuf:"bytes,19,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string         `protobuf:"bytes,20,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string         `protobuf:"bytes,21,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string         `protobuf:"bytes,22,opt,name=Website,proto3" json:"Website,omitempty"`
	Title      string         `protobuf:"bytes,23,opt,name=Title,proto3" json:"Title,omitempty"`
	Ratings    []*MovieRating `protobuf:"bytes,24,rep,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *MovieDetail) Reset() {
	*x = MovieDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetail) ProtoMessage() {}

func (x *MovieDetail) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetail.ProtoReflect.Descriptor instead.
func (*MovieDetail) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{6}
}

func (x *MovieDetail) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieDetail) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *MovieDetail) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *MovieDetail) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *MovieDetail) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *MovieDetail) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieDetail) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *MovieDetail) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *MovieDetail) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *MovieDetail) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MovieDetail) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MovieDetail) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *MovieDetail) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *MovieDetail) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *MovieDetail) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *MovieDetail) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *MovieDetail) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetail) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *MovieDetail) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *MovieDetail) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *MovieDetail) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *MovieDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetail) GetRatings() []*MovieRating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type MovieRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *MovieRating) Reset() {
	*x = MovieRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRating) ProtoMessage() {}

func (x *MovieRating) ProtoReflect() protoreflect.Message {
	mi := &file_handler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRating.ProtoReflect.Descriptor instead.
func (*MovieRating) Descriptor() ([]byte, []int) {
	return file_handler_proto_rawDescGZIP(), []int{7}
}

func (x *MovieRating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MovieRating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_handler_proto protoreflect.FileDescriptor

var file_handler_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x68, 0x0a, 0x0f,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x62, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x75, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0xfe, 0x04, 0x0a, 0x0b,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x6c, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6c, 0x6f,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x74, 0x61,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x56,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x3b, 0x0a, 0x0b,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x7b, 0x0a, 0x0c, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handler_proto_rawDescOnce sync.Once
	file_handler_proto_rawDescData = file_handler_proto_rawDesc
)

func file_handler_proto_rawDescGZIP() []byte {
	file_handler_proto_rawDescOnce.Do(func() {
		file_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_proto_rawDescData)
	})
	return file_handler_proto_rawDescData
}

var file_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_handler_proto_goTypes = []interface{}{
	(*IdRequest)(nil),       // 0: grpc.IdRequest
	(*GetRequest)(nil),      // 1: grpc.GetRequest
	(*MoviePagination)(nil), // 2: grpc.MoviePagination
	(*MovieResponse)(nil),   // 3: grpc.MovieResponse
	(*MovieMeta)(nil),       // 4: grpc.MovieMeta
	(*Movie)(nil),           // 5: grpc.Movie
	(*MovieDetail)(nil),     // 6: grpc.MovieDetail
	(*MovieRating)(nil),     // 7: grpc.MovieRating
}
var file_handler_proto_depIdxs = []int32{
	4, // 0: grpc.MoviePagination.data:type_name -> grpc.MovieMeta
	6, // 1: grpc.MovieResponse.data:type_name -> grpc.MovieDetail
	5, // 2: grpc.MovieMeta.list:type_name -> grpc.Movie
	7, // 3: grpc.MovieDetail.ratings:type_name -> grpc.MovieRating
	1, // 4: grpc.MovieService.GetMovie:input_type -> grpc.GetRequest
	0, // 5: grpc.MovieService.GetMovieDetail:input_type -> grpc.IdRequest
	2, // 6: grpc.MovieService.GetMovie:output_type -> grpc.MoviePagination
	3, // 7: grpc.MovieService.GetMovieDetail:output_type -> grpc.MovieResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_handler_proto_init() }
func file_handler_proto_init() {
	if File_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoviePagination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_proto_goTypes,
		DependencyIndexes: file_handler_proto_depIdxs,
		MessageInfos:      file_handler_proto_msgTypes,
	}.Build()
	File_handler_proto = out.File
	file_handler_proto_rawDesc = nil
	file_handler_proto_goTypes = nil
	file_handler_proto_depIdxs = nil
}
