package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopLikeDeleteAPIResponse 取消收藏 API返回值
// taobao.ailab.aicloud.top.like.delete
//
// 取消收藏
type TaobaoAilabAicloudTopLikeDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopLikeDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopLikeDeleteAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopLikeDeleteAPIResponseModel is 取消收藏 成功返回结果
type TaobaoAilabAicloudTopLikeDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_like_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 具体结果值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.IsSuccess = false
	m.Model = false
}

var poolTaobaoAilabAicloudTopLikeDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopLikeDeleteAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopLikeDeleteAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopLikeDeleteAPIResponse
func GetTaobaoAilabAicloudTopLikeDeleteAPIResponse() *TaobaoAilabAicloudTopLikeDeleteAPIResponse {
	return poolTaobaoAilabAicloudTopLikeDeleteAPIResponse.Get().(*TaobaoAilabAicloudTopLikeDeleteAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopLikeDeleteAPIResponse 将 TaobaoAilabAicloudTopLikeDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopLikeDeleteAPIResponse(v *TaobaoAilabAicloudTopLikeDeleteAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopLikeDeleteAPIResponse.Put(v)
}
