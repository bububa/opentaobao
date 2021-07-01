package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductGetAPIRequest
获得单个商品详情 API请求
alibaba.icbu.product.get

获取商品详情 */
type AlibabaIcbuProductGetAPIRequest struct {
	model.Params
	// 商品语种，目前只支持ENGLISH
	_language string
	// 混淆后的商品ID
	_productId string
}

// New
