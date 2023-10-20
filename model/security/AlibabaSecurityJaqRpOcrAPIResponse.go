package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpocrAPIResponse 聚安全实人认证证件OCR识别功能接口 API返回值
// alibaba.security.jaq.rp.ocr
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabasecurityjaqrpocrAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqrpocrAPIResponseModel
}

// AlibabasecurityjaqrpocrAPIResponseModel is 聚安全实人认证证件OCR识别功能接口 成功返回结果
type AlibabasecurityjaqrpocrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_ocr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Data *RpidCardBo `json:"data,omitempty" xml:"data,omitempty"`
}
