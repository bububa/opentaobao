package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopLikeListAPIResponse 列出收藏列表 API返回值
// taobao.ailab.aicloud.top.like.list
//
// 列出收藏列表
type TaobaoAilabAicloudTopLikeListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopLikeListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopLikeListAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopLikeListAPIResponseModel is 列出收藏列表 成功返回结果
type TaobaoAilabAicloudTopLikeListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_like_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopLikeListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopLikeListAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopLikeListAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopLikeListAPIResponse
func GetTaobaoAilabAicloudTopLikeListAPIResponse() *TaobaoAilabAicloudTopLikeListAPIResponse {
	return poolTaobaoAilabAicloudTopLikeListAPIResponse.Get().(*TaobaoAilabAicloudTopLikeListAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopLikeListAPIResponse 将 TaobaoAilabAicloudTopLikeListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopLikeListAPIResponse(v *TaobaoAilabAicloudTopLikeListAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopLikeListAPIResponse.Put(v)
}
