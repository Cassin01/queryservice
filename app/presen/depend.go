package presen

import (
	"queryservice/infra/gorm"
	"queryservice/presen/builder"
	"queryservice/presen/prepare"
	"queryservice/presen/server"

	"go.uber.org/fx"
)

var QueryDepend = fx.Options(
	gorm.RepDepend,
	fx.Provide( // プレゼンテーション層の依存定義
		builder.NewresultBuilderIMpl,
		server.NewcategoryServer,
		server.NewproductServerIMpl,
		prepare.NewQueryServer,
	),
	fx.Invoke(prepare.QueryServiceLifecycle), // 起動
)
