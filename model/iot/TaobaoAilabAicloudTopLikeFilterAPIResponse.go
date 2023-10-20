package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopLikeFilterAPIResponse 过滤列表歌曲存在于收藏列表的 API返回值
// taobao.ailab.aicloud.top.like.filter
//
// 过滤出传入列表歌曲存在于收藏列表的
type TaobaoAilabAicloudTopLikeFilterAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopLikeFilterAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeFilterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopLikeFilterAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopLikeFilterAPIResponseModel is 过滤列表歌曲存在于收藏列表的 成功返回结果
type TaobaoAilabAicloudTopLikeFilterAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_like_filter_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopLikeFilterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopLikeFilterAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopLikeFilterAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopLikeFilterAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopLikeFilterAPIResponse
func GetTaobaoAilabAicloudTopLikeFilterAPIResponse() *TaobaoAilabAicloudTopLikeFilterAPIResponse {
	return poolTaobaoAilabAicloudTopLikeFilterAPIResponse.Get().(*TaobaoAilabAicloudTopLikeFilterAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopLikeFilterAPIResponse 将 TaobaoAilabAicloudTopLikeFilterAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopLikeFilterAPIResponse(v *TaobaoAilabAicloudTopLikeFilterAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopLikeFilterAPIResponse.Put(v)
}
