package builder

import "github.com/Cassin01/samplepb/pb"

// 実行結果をXXXResult型に変換するインターフェース
type ResultBuilder interface {
	// *categories.Categoryを*pb.CategorResultに変換する
	BuildCategoryResult(source any) *pb.CategoryResult

	// []*categories.CAtegoryを*pb.CategoriesResultに変換する
	BuildCategoriesResult(source any) *pb.CategoriesResult

	// *prodcuts.Productを*pbProductResultに変換する
	BuildProductResult(source any) *pb.ProductResult

	// []*product.Productを*pb.ProductResultに変換する
	BuildProductResult(source any) *pb.ProductsResult

	// errs.CRUDERror, errs.InternalErrorを*pb.Errorに変換する
	BuildErrorResult(source any) *pb.Error
}
