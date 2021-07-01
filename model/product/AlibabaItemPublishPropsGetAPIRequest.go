package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemPublishPropsGetAPIRequest
商品级联属性信息获取 API请求
alibaba.item.publish.props.get

新商品发布，商品级联属性信息获取 */
type AlibabaItemPublishPropsGetAPIRequest struct {
	model.Params
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品类目ID
	_catId int64
	// 商品条码
	_barcode string
	// 类目属性渲染schema
	_schema string
	// 属性ID
	_propId int64
}

// New
