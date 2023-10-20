package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse 获取未读的消息数量 API返回值
// taobao.ailab.aicloud.top.message.get.unread.count
//
// 开放获取未读留言数量的接口
type TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponseModel is 获取未读的消息数量 成功返回结果
type TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_message_get_unread_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 未读留言的数量
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.Model = 0
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse
func GetTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse() *TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse {
	return poolTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse.Get().(*TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse 将 TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse(v *TaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageGetUnreadCountAPIResponse.Put(v)
}
