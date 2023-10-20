package alihealthmdeer

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerScienceDeletearticleAPIResponse 文章删除 API返回值
// alibaba.alihealth.mdeer.science.deletearticle
//
// 三方同步文章删除
type AlibabaAlihealthMdeerScienceDeletearticleAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel
}

// AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel is 文章删除 成功返回结果
type AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_mdeer_science_deletearticle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否删除成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
