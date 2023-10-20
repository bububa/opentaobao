package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqresourcefetchAPIResponse 获取资源文件 API返回值
// alibaba.security.jaq.resource.fetch
//
// 在前向化验证流程中提供资源文件服务
type AlibabasecurityjaqresourcefetchAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqresourcefetchAPIResponseModel
}

// AlibabasecurityjaqresourcefetchAPIResponseModel is 获取资源文件 成功返回结果
type AlibabasecurityjaqresourcefetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_resource_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取资源结果
	Data *JaqResourceResult `json:"data,omitempty" xml:"data,omitempty"`
}
