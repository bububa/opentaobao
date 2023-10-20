package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimAppChatlogsGetAPIResponse openim应用聊天记录查询 API返回值
// taobao.openim.app.chatlogs.get
//
// 查询openim应用的聊天记录
type TaobaoOpenimAppChatlogsGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimAppChatlogsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimAppChatlogsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimAppChatlogsGetAPIResponseModel).Reset()
}

// TaobaoOpenimAppChatlogsGetAPIResponseModel is openim应用聊天记录查询 成功返回结果
type TaobaoOpenimAppChatlogsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_app_chatlogs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *EsMessageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimAppChatlogsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOpenimAppChatlogsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimAppChatlogsGetAPIResponse)
	},
}

// GetTaobaoOpenimAppChatlogsGetAPIResponse 从 sync.Pool 获取 TaobaoOpenimAppChatlogsGetAPIResponse
func GetTaobaoOpenimAppChatlogsGetAPIResponse() *TaobaoOpenimAppChatlogsGetAPIResponse {
	return poolTaobaoOpenimAppChatlogsGetAPIResponse.Get().(*TaobaoOpenimAppChatlogsGetAPIResponse)
}

// ReleaseTaobaoOpenimAppChatlogsGetAPIResponse 将 TaobaoOpenimAppChatlogsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimAppChatlogsGetAPIResponse(v *TaobaoOpenimAppChatlogsGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenimAppChatlogsGetAPIResponse.Put(v)
}
