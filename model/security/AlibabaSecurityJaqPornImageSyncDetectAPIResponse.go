package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqPornImageSyncDetectAPIResponse 聚安全智能鉴黄同步检测接口 API返回值
// alibaba.security.jaq.porn.image.sync.detect
//
// 同步黄图图像检测接口
type AlibabaSecurityJaqPornImageSyncDetectAPIResponse struct {
	model.CommonResponse
	AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel
}

// AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel is 聚安全智能鉴黄同步检测接口 成功返回结果
type AlibabaSecurityJaqPornImageSyncDetectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_porn_image_sync_detect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参结构体
	Data *JaqPornImageDetectResult `json:"data,omitempty" xml:"data,omitempty"`
}
