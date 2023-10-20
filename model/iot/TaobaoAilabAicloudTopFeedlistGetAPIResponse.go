package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopFeedlistGetAPIResponse 获取对话流列表 API返回值
// taobao.ailab.aicloud.top.feedlist.get
//
// 获取指定应用的对话流信息
type TaobaoAilabAicloudTopFeedlistGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopFeedlistGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopFeedlistGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopFeedlistGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopFeedlistGetAPIResponseModel is 获取对话流列表 成功返回结果
type TaobaoAilabAicloudTopFeedlistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_feedlist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopFeedlistGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopFeedlistGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopFeedlistGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopFeedlistGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopFeedlistGetAPIResponse
func GetTaobaoAilabAicloudTopFeedlistGetAPIResponse() *TaobaoAilabAicloudTopFeedlistGetAPIResponse {
	return poolTaobaoAilabAicloudTopFeedlistGetAPIResponse.Get().(*TaobaoAilabAicloudTopFeedlistGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopFeedlistGetAPIResponse 将 TaobaoAilabAicloudTopFeedlistGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopFeedlistGetAPIResponse(v *TaobaoAilabAicloudTopFeedlistGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopFeedlistGetAPIResponse.Put(v)
}
