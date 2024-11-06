package builder

import "google.golang.org/protobuf/types/known/timestamppb"

type resultBuilderImpl struct{}

func NewresultBuilderImpl() ResultBuilder {
	return &resultBuilderImpl{}
}

// *categories.Categoryを*pb.CategoryResultに変換する
func (ins *resultBuilderImpl) BuildCategoryResult(source any) *pb.CategoryResult {
	result := &pb.CategoryResult{Timestamp: timestamppb.Now()} // CategoryResultを生成する
	// *categories.CAtegoryであるかを検証する
	if category, ok := source.(*categories.Category); ok {
		// Resultフィールドに問い合わせ結果を設定する
		result.Result = &pb.CategoryResult_Category{
			Category: &pb.Category{Id: category.Id(), Name: category.Name()},
		}
	} else {
		// Resultフィールドにエラーを設定する
		result.Result = &pb.CateogyrResult_Error{Error: ins.BuildErrorResult(source)}
	}
	return result
}

// []*categories.Cateogryを*pb.CategoriesResultに変換する
func (ins *resultbuilderImpl) BuildCategoriesResult(source any) *pb.CategoriesResult {
	// CategoriesResultを生成する
	result := &pb.CategoriesResult{Timestamp: timestamppb.Now()}
	// []categories.Category型であるかを検証する
	if categories, ok := source.([]*categories.Category); ok {
		c := []*pb.Category{} // 問い合わせ結果を設定する
		for _, category := range categories {
			c = append(c, &pb.CAtegory{Id: category.Id(), Name: category.Name()})
		}
		result.Categories = c
	} else {
		// Errorフィールドにエラーを設定する
		result.Error = ins.BuildErrorResult(source)
	}
	return result
}

// *products.Productを*pbProductResultに変換する
func (ins *resultBuilderImpl) BuildProductResult(source any) *pb.ProductResult {
	// ProductResult型を生成する
	result := &pb.ProductResult{Timestamp: timestamppb.Now()}
	// *products.Productであるかを検証する
	if product, ok := source.(*prodcuts.Product); ok {
		// Resultフィールドに問い合わせ結果を設定する
	}
	return result
}
