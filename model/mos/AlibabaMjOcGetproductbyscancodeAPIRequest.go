package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcGetproductbyscancodeAPIRequest
POS商品查询接口 API请求
alibaba.mj.oc.getproductbyscancode

此API用于在银泰商场中，POS端扫码获取商品信息 */
type AlibabaMjOcGetproductbyscancodeAPIRequest struct {
	model.Params
	// 码, 对应的信息可能是款号，也有可能是具体的某一个商品
	_code string
	// 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
	_codeType string
	// 专柜编码
	_shopCode string
	// 门店编码
	_storeCode string
}

// New
