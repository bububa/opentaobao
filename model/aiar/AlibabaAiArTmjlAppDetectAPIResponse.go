package aiar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaiartmjlappdetectAPIResponse 天猫精灵扫一扫入口的服务 API返回值
// alibaba.ai.ar.tmjl.app.detect
//
// 天猫精灵扫一扫入口的图像检测服务
type AlibabaaiartmjlappdetectAPIResponse struct {
	model.CommonResponse
	AlibabaaiartmjlappdetectAPIResponseModel
}

// AlibabaaiartmjlappdetectAPIResponseModel is 天猫精灵扫一扫入口的服务 成功返回结果
type AlibabaaiartmjlappdetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_ar_tmjl_app_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Results string `json:"results,omitempty" xml:"results,omitempty"`
}
