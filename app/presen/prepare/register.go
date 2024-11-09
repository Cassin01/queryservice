package prepare

import (
	"github.com/Cassin01/samplepb/pb"
	"google.golang.org/grpc"
)

// gRPCサーバの生成とQueryServiceの登録
type QueryServer struct {
	Server *grpc.Server // gRPCServer
}

// コンストラクタ
func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	// gRPCサーバを生成する(インターセプタの追加)
	server := grpc.NewServer()

	// CategoryQueryServerを登録する
	pb.RegisterCategoryQueryServer(server, category)
	// ProductQueryServerを登録する
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{Server: server}
}
