package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ccodeconvertAPIResponse b2c转码 API返回值
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
type Alibabanlifeb2ccodeconvertAPIResponse struct {
	model.CommonResponse
	Alibabanlifeb2ccodeconvertAPIResponseModel
}

// Alibabanlifeb2ccodeconvertAPIResponseModel is b2c转码 成功返回结果
type Alibabanlifeb2ccodeconvertAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_code_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *ItemCodeConvertResponse `json:"data,omitempty" xml:"data,omitempty"`
}
