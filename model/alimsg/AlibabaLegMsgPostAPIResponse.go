package alimsg

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegMsgPostAPIResponse 集团法务消息发送 API返回值
// alibaba.leg.msg.post
//
// 消息发送能力
type AlibabaLegMsgPostAPIResponse struct {
	model.CommonResponse
	AlibabaLegMsgPostAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegMsgPostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegMsgPostAPIResponseModel).Reset()
}

// AlibabaLegMsgPostAPIResponseModel is 集团法务消息发送 成功返回结果
type AlibabaLegMsgPostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_leg_msg_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegMsgPostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegMsgPostAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegMsgPostAPIResponse)
	},
}

// GetAlibabaLegMsgPostAPIResponse 从 sync.Pool 获取 AlibabaLegMsgPostAPIResponse
func GetAlibabaLegMsgPostAPIResponse() *AlibabaLegMsgPostAPIResponse {
	return poolAlibabaLegMsgPostAPIResponse.Get().(*AlibabaLegMsgPostAPIResponse)
}

// ReleaseAlibabaLegMsgPostAPIResponse 将 AlibabaLegMsgPostAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegMsgPostAPIResponse(v *AlibabaLegMsgPostAPIResponse) {
	v.Reset()
	poolAlibabaLegMsgPostAPIResponse.Put(v)
}
