package prepare

import (
	"context"

	"queryservice/presen/interceptor"

	"github.com/Cassin01/samplepb/pb"
	"google.golang.org/grpc"
)

// gRPCサーバの生成とQueryServiceの登録
type QueryServer struct {
	Server *grpc.Server // gRPCServer
}

// インターセプタをチェーン化して実行する
func chainUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// LoggingInterceptor内部でUUIDValidationInterceptorと
	// handlerを呼び出す新しいhandlerを作成する
	newHandler := func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
		// UUID形式の検証インターセプタを実行する
		return interceptor.UUIDValidationInterceptor(currentCtx, currentReq, info, handler)
	}
	// ログ出力インターセプタを実行する
	return interceptor.LoggingInterceptor(ctx, req, info, newHandler)
}

// コンストラクタ
func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	// インターセプタの追加
	serverOpts := []grpc.ServerOption{
		// ログ出力、入力検証インターセプタを登録インターセプタを登録
		grpc.UnaryInterceptor(chainUnaryInterceptor),
	}

	// gRPCサーバを生成する(インターセプタの追加)
	server := grpc.NewServer(serverOpts...)

	// CategoryQueryServerを登録する
	pb.RegisterCategoryQueryServer(server, category)
	// ProductQueryServerを登録する
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{Server: server}
}
