// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: modestyMart_submodule/products.proto

package genproto

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

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string      `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Category    *Categories `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Stock       int32       `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
	CreatedAt   string      `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{0}
}

func (x *Products) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Products) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Products) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Products) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Products) GetCategory() *Categories {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Products) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Products) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Category    string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Stock       int32  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *CreateProductReq) Reset() {
	*x = CreateProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReq) ProtoMessage() {}

func (x *CreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReq.ProtoReflect.Descriptor instead.
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductReq) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateProductReq) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductReq) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type UpdateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body *UpdateProduct `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *UpdateProductReq) Reset() {
	*x = UpdateProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductReq) ProtoMessage() {}

func (x *UpdateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductReq.ProtoReflect.Descriptor instead.
func (*UpdateProductReq) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProductReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductReq) GetBody() *UpdateProduct {
	if x != nil {
		return x.Body
	}
	return nil
}

type UpdateProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Category    string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Stock       int32  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *UpdateProduct) Reset() {
	*x = UpdateProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProduct) ProtoMessage() {}

func (x *UpdateProduct) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProduct.ProtoReflect.Descriptor instead.
func (*UpdateProduct) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProduct) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *UpdateProduct) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateProduct) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type ListAllProductsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price      string      `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Stock      int32       `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *ListAllProductsReq) Reset() {
	*x = ListAllProductsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllProductsReq) ProtoMessage() {}

func (x *ListAllProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllProductsReq.ProtoReflect.Descriptor instead.
func (*ListAllProductsReq) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllProductsReq) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListAllProductsReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListAllProductsReq) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ListAllProductsReq) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type ListAllProductsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products *Products `protobuf:"bytes,1,opt,name=products,proto3" json:"products,omitempty"`
	Count    int32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListAllProductsRes) Reset() {
	*x = ListAllProductsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllProductsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllProductsRes) ProtoMessage() {}

func (x *ListAllProductsRes) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllProductsRes.ProtoReflect.Descriptor instead.
func (*ListAllProductsRes) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{5}
}

func (x *ListAllProductsRes) GetProducts() *Products {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ListAllProductsRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID string `protobuf:"bytes,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
}

func (x *GetCategoryReq) Reset() {
	*x = GetCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryReq) ProtoMessage() {}

func (x *GetCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryReq.ProtoReflect.Descriptor instead.
func (*GetCategoryReq) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{6}
}

func (x *GetCategoryReq) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

type GetCategoryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products *Products `protobuf:"bytes,1,opt,name=products,proto3" json:"products,omitempty"`
	Count    int32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetCategoryRes) Reset() {
	*x = GetCategoryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRes) ProtoMessage() {}

func (x *GetCategoryRes) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRes.ProtoReflect.Descriptor instead.
func (*GetCategoryRes) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{7}
}

func (x *GetCategoryRes) GetProducts() *Products {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetCategoryRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ProductsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	CategoryID  string `protobuf:"bytes,5,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
	Stock       int32  `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ProductsRes) Reset() {
	*x = ProductsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsRes) ProtoMessage() {}

func (x *ProductsRes) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsRes.ProtoReflect.Descriptor instead.
func (*ProductsRes) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{8}
}

func (x *ProductsRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductsRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductsRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductsRes) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ProductsRes) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *ProductsRes) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductsRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetProductsByPriceRangeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinPrice string `protobuf:"bytes,1,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice string `protobuf:"bytes,2,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
}

func (x *GetProductsByPriceRangeReq) Reset() {
	*x = GetProductsByPriceRangeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modestyMart_submodule_products_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsByPriceRangeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsByPriceRangeReq) ProtoMessage() {}

func (x *GetProductsByPriceRangeReq) ProtoReflect() protoreflect.Message {
	mi := &file_modestyMart_submodule_products_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsByPriceRangeReq.ProtoReflect.Descriptor instead.
func (*GetProductsByPriceRangeReq) Descriptor() ([]byte, []int) {
	return file_modestyMart_submodule_products_proto_rawDescGZIP(), []int{9}
}

func (x *GetProductsByPriceRangeReq) GetMinPrice() string {
	if x != nil {
		return x.MinPrice
	}
	return ""
}

func (x *GetProductsByPriceRangeReq) GetMaxPrice() string {
	if x != nil {
		return x.MaxPrice
	}
	return ""
}

var File_modestyMart_submodule_products_proto protoreflect.FileDescriptor

var file_modestyMart_submodule_products_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x79, 0x4d, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d,
	0x6f, 0x64, 0x65, 0x73, 0x74, 0x79, 0x4d, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x79, 0x4d, 0x61, 0x72, 0x74, 0x5f, 0x73,
	0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x4c, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x31,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x22, 0x57, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x2b,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0xbe, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x56, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0xbb, 0x03, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12,
	0x57, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modestyMart_submodule_products_proto_rawDescOnce sync.Once
	file_modestyMart_submodule_products_proto_rawDescData = file_modestyMart_submodule_products_proto_rawDesc
)

func file_modestyMart_submodule_products_proto_rawDescGZIP() []byte {
	file_modestyMart_submodule_products_proto_rawDescOnce.Do(func() {
		file_modestyMart_submodule_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_modestyMart_submodule_products_proto_rawDescData)
	})
	return file_modestyMart_submodule_products_proto_rawDescData
}

var file_modestyMart_submodule_products_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_modestyMart_submodule_products_proto_goTypes = []any{
	(*Products)(nil),                   // 0: proto.Products
	(*CreateProductReq)(nil),           // 1: proto.CreateProductReq
	(*UpdateProductReq)(nil),           // 2: proto.UpdateProductReq
	(*UpdateProduct)(nil),              // 3: proto.UpdateProduct
	(*ListAllProductsReq)(nil),         // 4: proto.ListAllProductsReq
	(*ListAllProductsRes)(nil),         // 5: proto.ListAllProductsRes
	(*GetCategoryReq)(nil),             // 6: proto.GetCategoryReq
	(*GetCategoryRes)(nil),             // 7: proto.GetCategoryRes
	(*ProductsRes)(nil),                // 8: proto.ProductsRes
	(*GetProductsByPriceRangeReq)(nil), // 9: proto.GetProductsByPriceRangeReq
	(*Categories)(nil),                 // 10: proto.Categories
	(*Pagination)(nil),                 // 11: proto.Pagination
	(*GetById)(nil),                    // 12: proto.GetById
	(*Void)(nil),                       // 13: proto.Void
}
var file_modestyMart_submodule_products_proto_depIdxs = []int32{
	10, // 0: proto.Products.category:type_name -> proto.Categories
	3,  // 1: proto.UpdateProductReq.body:type_name -> proto.UpdateProduct
	11, // 2: proto.ListAllProductsReq.pagination:type_name -> proto.Pagination
	0,  // 3: proto.ListAllProductsRes.products:type_name -> proto.Products
	0,  // 4: proto.GetCategoryRes.products:type_name -> proto.Products
	1,  // 5: proto.ProductsService.CreateProduct:input_type -> proto.CreateProductReq
	2,  // 6: proto.ProductsService.UpdateProduct:input_type -> proto.UpdateProductReq
	12, // 7: proto.ProductsService.GetProduct:input_type -> proto.GetById
	12, // 8: proto.ProductsService.DeleteProduct:input_type -> proto.GetById
	4,  // 9: proto.ProductsService.ListAllProducts:input_type -> proto.ListAllProductsReq
	6,  // 10: proto.ProductsService.GetCategory:input_type -> proto.GetCategoryReq
	9,  // 11: proto.ProductsService.GetProductsByPriceRange:input_type -> proto.GetProductsByPriceRangeReq
	13, // 12: proto.ProductsService.CreateProduct:output_type -> proto.Void
	13, // 13: proto.ProductsService.UpdateProduct:output_type -> proto.Void
	0,  // 14: proto.ProductsService.GetProduct:output_type -> proto.Products
	13, // 15: proto.ProductsService.DeleteProduct:output_type -> proto.Void
	5,  // 16: proto.ProductsService.ListAllProducts:output_type -> proto.ListAllProductsRes
	7,  // 17: proto.ProductsService.GetCategory:output_type -> proto.GetCategoryRes
	5,  // 18: proto.ProductsService.GetProductsByPriceRange:output_type -> proto.ListAllProductsRes
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_modestyMart_submodule_products_proto_init() }
func file_modestyMart_submodule_products_proto_init() {
	if File_modestyMart_submodule_products_proto != nil {
		return
	}
	file_modestyMart_submodule_common_proto_init()
	file_modestyMart_submodule_categories_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_modestyMart_submodule_products_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Products); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProductReq); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProductReq); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProduct); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllProductsReq); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllProductsRes); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCategoryReq); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetCategoryRes); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ProductsRes); i {
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
		file_modestyMart_submodule_products_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetProductsByPriceRangeReq); i {
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
			RawDescriptor: file_modestyMart_submodule_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modestyMart_submodule_products_proto_goTypes,
		DependencyIndexes: file_modestyMart_submodule_products_proto_depIdxs,
		MessageInfos:      file_modestyMart_submodule_products_proto_msgTypes,
	}.Build()
	File_modestyMart_submodule_products_proto = out.File
	file_modestyMart_submodule_products_proto_rawDesc = nil
	file_modestyMart_submodule_products_proto_goTypes = nil
	file_modestyMart_submodule_products_proto_depIdxs = nil
}
