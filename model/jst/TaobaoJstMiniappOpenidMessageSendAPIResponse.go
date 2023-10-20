package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstMiniappOpenidMessageSendAPIResponse 单个openId用户短信发送 API返回值
// taobao.jst.miniapp.openid.message.send
//
// 单个openId用户短信发送
type TaobaoJstMiniappOpenidMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoJstMiniappOpenidMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstMiniappOpenidMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstMiniappOpenidMessageSendAPIResponseModel).Reset()
}

// TaobaoJstMiniappOpenidMessageSendAPIResponseModel is 单个openId用户短信发送 成功返回结果
type TaobaoJstMiniappOpenidMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_miniapp_openid_message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短信回执码
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 请求失败错误信息
	RErrMsg string `json:"r_err_msg,omitempty" xml:"r_err_msg,omitempty"`
	// 请求code
	RCode int64 `json:"r_code,omitempty" xml:"r_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstMiniappOpenidMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.RErrMsg = ""
	m.RCode = 0
}

var poolTaobaoJstMiniappOpenidMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstMiniappOpenidMessageSendAPIResponse)
	},
}

// GetTaobaoJstMiniappOpenidMessageSendAPIResponse 从 sync.Pool 获取 TaobaoJstMiniappOpenidMessageSendAPIResponse
func GetTaobaoJstMiniappOpenidMessageSendAPIResponse() *TaobaoJstMiniappOpenidMessageSendAPIResponse {
	return poolTaobaoJstMiniappOpenidMessageSendAPIResponse.Get().(*TaobaoJstMiniappOpenidMessageSendAPIResponse)
}

// ReleaseTaobaoJstMiniappOpenidMessageSendAPIResponse 将 TaobaoJstMiniappOpenidMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstMiniappOpenidMessageSendAPIResponse(v *TaobaoJstMiniappOpenidMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoJstMiniappOpenidMessageSendAPIResponse.Put(v)
}
