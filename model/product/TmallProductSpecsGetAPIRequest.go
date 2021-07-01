package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSpecsGetAPIRequest
获取产品的规格信息 API请求
tmall.product.specs.get

按product_id或品牌下载产品规格，返回一组的产品规格信息。 */
type TmallProductSpecsGetAPIRequest struct {
	model.Params
	// 产品的ID。这个不能和properties和cat_id同时起效果<br>properties 和cat_id 均不传时，该参数必传。
	_productId int64
	// 关键属性的字符串，pid:vid;pid:vid该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。<br>product_id 不传时该参数必传
	_properties string
	// 类目的ID号，该id必须和properties同时传入。而且只有当product_id不传入的时候才起效果。<br> product_id不传时，该参数必传
	_catId int64
}

// New
