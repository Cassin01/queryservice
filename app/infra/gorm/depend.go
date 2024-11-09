package gorm

import (
	"queryservice/infra/gorm/handler"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// データベース接続
var DBModule = fx.Provide(func() (*gorm.DB, error) {
	return handler.ConnectDB()
})

// gorm.DB, Adapter, Repositoryの依存定義
var RepDepend = fx.Options(
	DBModule,
	fx.Provide(
		// Adapterインターフェイス実装のコンストラクタを指定
		adapter.NewcateogryAdapterImpl,
		adapter.NewproductAdapterImpl,

		// Repositoryインターフェイス実装のコンストラクタを指定
		repository.NewCategoryRepositoryGORM,
		repository.NewProductRepositoryGORM,
	),
)
