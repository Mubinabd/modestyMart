package storage

import (
	pb "github.com/Mubinabd/modestyMart/internal/pkg/genproto"
)

type StorageI interface {
	Product() ProductI
	Order() OrderI
	Payment() PaymentI
	Category() CategoryI
	Cart() CartI
	Auth() AuthI
	User() UserI
	Notification() NotificationI
}

type ProductI interface {
	CreateProduct(req *pb.CreateProductReq) (*pb.Void, error)
	UpdateProduct(req *pb.UpdateProductReq) (*pb.Void, error)
	GetProduct(req *pb.GetById) (*pb.Products, error)
	DeleteProduct(req *pb.GetById) (*pb.Void, error)
	ListAllProducts(req *pb.ListAllProductsReq) (*pb.ListAllProductsRes, error)
	GetCategory(req *pb.GetCategoryReq) (*pb.GetCategoryRes, error)
	GetProductsByPriceRange(req *pb.GetProductsByPriceRangeReq) (*pb.ListAllProductsRes, error)
}

type OrderI interface {
	CreateOrder(req *pb.CreateOrderReq) (*pb.Void, error)
	UpdateOrder(req *pb.UpdateOrderReq) (*pb.Void, error)
	GetOrder(req *pb.GetById) (*pb.Orders, error)
	DeleteOrder(req *pb.GetById) (*pb.Void, error)
	ListAllOrders(req *pb.ListOrdersReq) (*pb.ListOrdersRes, error)
	GetOrderByProductID(req *pb.OrderByProductId) (*pb.GetOrdersRes, error)
}
type PaymentI interface {
	CreatePayment(req *pb.CreatePaymentReq) (*pb.Void, error)
	GetPayment(req *pb.GetById) (*pb.Payment, error)
	ListPayments(req *pb.ListPaymentsReq) (*pb.ListPaymentsRes, error)
}

type CategoryI interface {
	CreateCategory(req *pb.CreateCategoryReq) (*pb.Void, error)
	ListCategories(req *pb.ListAllCategoriesReq) (*pb.ListCategoriesRes, error)
	UpdateCategory(req *pb.UpdateCategoryReq) (*pb.Void, error)
	DeleteCategory(req *pb.GetById) (*pb.Void, error)
	GetCategory(req *pb.GetById) (*pb.Categories, error)
}

type AuthI interface {
	Register(*pb.RegisterReq) (*pb.Void, error)
	Login(*pb.LoginReq) (*pb.User, error)
	ForgotPassword(*pb.GetByEmail) (*pb.Void, error)
	ResetPassword(*pb.ResetPassReq) (*pb.Void, error)
	SaveRefreshToken(*pb.RefToken) (*pb.Void, error)
	GetAllUsers(*pb.ListUserReq) (*pb.ListUserRes, error)
}

type UserI interface {
	GetProfile(*pb.GetById) (*pb.UserRes, error)
	EditProfile(*pb.UserRes) (*pb.UserRes, error)
	ChangePassword(*pb.ChangePasswordReq) (*pb.Void, error)
	GetSetting(*pb.GetById) (*pb.Setting, error)
	EditSetting(*pb.SettingReq) (*pb.Void, error)
	DeleteUser(*pb.GetById) (*pb.Void, error)
}

type CartI interface {
	CreateCart(req *pb.CreateCartReq) (*pb.Void, error)
	GetCart(req *pb.GetById) (*pb.Cart, error)
	ListAllCarts(req *pb.ListAllCartReq) (*pb.ListAllCartRes, error)
}	

type NotificationI interface {
	CreateNotification(req *pb.NotificationCreate) (*pb.Void, error)
	NotifyAll(req *pb.NotificationMessage) (*pb.Void, error)
	DeleteNotificcation(id *pb.GetById) (*pb.Void, error)
	UpdateNotificcation(req *pb.NotificationUpdate) (*pb.Void, error)
	GetNotifications(req *pb.NotifFilter) (*pb.NotificationList, error)
	GetNotification(id *pb.GetById) (*pb.NotificationGet, error)
}
