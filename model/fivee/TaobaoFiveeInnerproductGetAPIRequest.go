package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFiveeInnerproductGetAPIRequest
国产商品资质查询 API请求
taobao.fivee.innerproduct.get

资质共享平台，国产商品查询 */
type TaobaoFiveeInnerproductGetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 条形码
	_paramBarcode string
}

// New
