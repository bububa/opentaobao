package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cCodeConvertAPIRequest
b2c转码 API请求
alibaba.nlife.b2c.code.convert

将商品的URL转码，ISV将该码写入RFID */
type AlibabaNlifeB2cCodeConvertAPIRequest struct {
	model.Params
	// 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
	_storeId string
	// 商品URL
	_url string
}

// New
