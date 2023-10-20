package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse 儿童音频列表 API返回值
// taobao.ailab.aicloud.top.freelisten.childrenalbum
//
// 儿童音频列表
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel is 儿童音频列表 成功返回结果
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_freelisten_childrenalbum_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse
func GetTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse() *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse {
	return poolTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse.Get().(*TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse 将 TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse(v *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopFreelistenChildrenalbumAPIResponse.Put(v)
}
