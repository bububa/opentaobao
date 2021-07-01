package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFiveeImportproductGetAPIRequest
进口商品查询 API请求
taobao.fivee.importproduct.get

资质共享平台查询进口商品信息 */
type TaobaoFiveeImportproductGetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBuCode string
	// 条形码
	_paramBarcode string
}

// New
