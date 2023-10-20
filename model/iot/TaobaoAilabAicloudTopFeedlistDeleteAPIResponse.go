package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopFeedlistDeleteAPIResponse 删除单条对话流信息 API返回值
// taobao.ailab.aicloud.top.feedlist.delete
//
// 删除指定的某一条对话流信息
type TaobaoAilabAicloudTopFeedlistDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopFeedlistDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopFeedlistDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopFeedlistDeleteAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopFeedlistDeleteAPIResponseModel is 删除单条对话流信息 成功返回结果
type TaobaoAilabAicloudTopFeedlistDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_feedlist_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// success
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopFeedlistDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.Model = ""
	m.IsSuccess = ""
}

var poolTaobaoAilabAicloudTopFeedlistDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopFeedlistDeleteAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopFeedlistDeleteAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopFeedlistDeleteAPIResponse
func GetTaobaoAilabAicloudTopFeedlistDeleteAPIResponse() *TaobaoAilabAicloudTopFeedlistDeleteAPIResponse {
	return poolTaobaoAilabAicloudTopFeedlistDeleteAPIResponse.Get().(*TaobaoAilabAicloudTopFeedlistDeleteAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopFeedlistDeleteAPIResponse 将 TaobaoAilabAicloudTopFeedlistDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopFeedlistDeleteAPIResponse(v *TaobaoAilabAicloudTopFeedlistDeleteAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopFeedlistDeleteAPIResponse.Put(v)
}
