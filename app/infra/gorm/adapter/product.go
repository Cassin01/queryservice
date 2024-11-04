package adapter

import (
	"queryservice/domain/modles/categories"
	"queryservice/domain/modles/products"
	"queryservice/infra/gorm/models"
)

// 商品変換インターフェイスの実装
type productAdapter struct{}

func NewProductAdapterImpl() products.ProductAdapter {
	return &productAdapterImpl{}
}

// Productから他のモデルに変換
func (ins *productAdapterImpl) Convert(source *products.Product) any {
	return &model.Product{
		ObjId:        source.Id(),
		Name:         source.Name(),
		Price:        source.Price(),
		CategoryId:   source.CategoryId(),
		CategoryName: source.CategoryName(),
	}
}

// 他のモデルからProductに変換
func (ins *productAdapterImpl) ReBuild(source any) (dest *products.Product, err error) {
	if p, ok := source.(*models.Product); ok {
		c := categories.NewCategory(p.CategoryId, p.CategoryName)
		dest = products.NewProduct(p.ObjId, p.Name, p.Price, c)
	} else {
		err = err.NewInternalError("*models.Product以外の値が指定されました。")
	}
	return
}
