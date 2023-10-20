package nlife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cCodeConvertAPIResponse b2c转码 API返回值
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
type AlibabaNlifeB2cCodeConvertAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cCodeConvertAPIResponseModel
}

// AlibabaNlifeB2cCodeConvertAPIResponseModel is b2c转码 成功返回结果
type AlibabaNlifeB2cCodeConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_code_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *ItemCodeConvertResponse `json:"data,omitempty" xml:"data,omitempty"`
}
