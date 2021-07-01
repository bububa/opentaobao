package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFiveeInnerproductPublishAPIRequest
国产商品发布 API请求
taobao.fivee.innerproduct.publish

资质共享平台国产商品发布 */
type TaobaoFiveeInnerproductPublishAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 国产商品
	_paramInnerProduct *InnerProduct
}

// New
