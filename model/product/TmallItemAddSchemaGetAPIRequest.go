package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemAddSchemaGetAPIRequest
天猫发布商品规则获取 API请求
tmall.item.add.schema.get

通过类目以及productId获取商品发布规则； */
type TmallItemAddSchemaGetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 商品发布的目标product_id
	_productId int64
	// 发布商品类型，一口价填“b”，拍卖填"a"
	_type string
	// 正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
	_isvInit bool
}

// New
