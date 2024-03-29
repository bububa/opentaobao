package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesConfirmAPIResponse 确认消费消息的状态 API返回值
// taobao.tmc.messages.confirm
//
// 确认消费消息的状态
type TaobaoTmcMessagesConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTmcMessagesConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcMessagesConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcMessagesConfirmAPIResponseModel).Reset()
}

// TaobaoTmcMessagesConfirmAPIResponseModel is 确认消费消息的状态 成功返回结果
type TaobaoTmcMessagesConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_messages_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcMessagesConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTmcMessagesConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcMessagesConfirmAPIResponse)
	},
}

// GetTaobaoTmcMessagesConfirmAPIResponse 从 sync.Pool 获取 TaobaoTmcMessagesConfirmAPIResponse
func GetTaobaoTmcMessagesConfirmAPIResponse() *TaobaoTmcMessagesConfirmAPIResponse {
	return poolTaobaoTmcMessagesConfirmAPIResponse.Get().(*TaobaoTmcMessagesConfirmAPIResponse)
}

// ReleaseTaobaoTmcMessagesConfirmAPIResponse 将 TaobaoTmcMessagesConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcMessagesConfirmAPIResponse(v *TaobaoTmcMessagesConfirmAPIResponse) {
	v.Reset()
	poolTaobaoTmcMessagesConfirmAPIResponse.Put(v)
}
