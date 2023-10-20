package ju

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityWechatLoginAPIResponse 聚划算用增淘外社群登录 API返回值
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
type AlibabaJhsCommunityWechatLoginAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityWechatLoginAPIResponseModel
}

// AlibabaJhsCommunityWechatLoginAPIResponseModel is 聚划算用增淘外社群登录 成功返回结果
type AlibabaJhsCommunityWechatLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_wechat_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 登录会话
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
}
