package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopLikeAddAPIResponse 增加收藏 API返回值
// taobao.ailab.aicloud.top.like.add
//
// 将制定内容加入收藏
type TaobaoAilabAicloudTopLikeAddAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopLikeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopLikeAddAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopLikeAddAPIResponseModel is 增加收藏 成功返回结果
type TaobaoAilabAicloudTopLikeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_like_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 标志是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 具体的结果值
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.IsSuccess = false
	m.Model = false
}

var poolTaobaoAilabAicloudTopLikeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopLikeAddAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopLikeAddAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopLikeAddAPIResponse
func GetTaobaoAilabAicloudTopLikeAddAPIResponse() *TaobaoAilabAicloudTopLikeAddAPIResponse {
	return poolTaobaoAilabAicloudTopLikeAddAPIResponse.Get().(*TaobaoAilabAicloudTopLikeAddAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopLikeAddAPIResponse 将 TaobaoAilabAicloudTopLikeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopLikeAddAPIResponse(v *TaobaoAilabAicloudTopLikeAddAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopLikeAddAPIResponse.Put(v)
}
