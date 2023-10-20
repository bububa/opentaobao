package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudocrcheckAPIResponse ocr同时实名校验 API返回值
// alibaba.security.jaq.rp.cloud.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabasecurityjaqrpcloudocrcheckAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqrpcloudocrcheckAPIResponseModel
}

// AlibabasecurityjaqrpcloudocrcheckAPIResponseModel is ocr同时实名校验 成功返回结果
type AlibabasecurityjaqrpcloudocrcheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_cloud_ocr_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Data *RpidCard `json:"data,omitempty" xml:"data,omitempty"`
}
