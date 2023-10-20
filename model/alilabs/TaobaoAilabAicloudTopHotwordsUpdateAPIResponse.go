package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopHotwordsUpdateAPIResponse 更新热词 API返回值
// taobao.ailab.aicloud.top.hotwords.update
//
// 更新ASR热词
type TaobaoAilabAicloudTopHotwordsUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopHotwordsUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopHotwordsUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopHotwordsUpdateAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopHotwordsUpdateAPIResponseModel is 更新热词 成功返回结果
type TaobaoAilabAicloudTopHotwordsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_hotwords_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// baseresult
	Baseresult *BaseResult `json:"baseresult,omitempty" xml:"baseresult,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopHotwordsUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Baseresult = nil
}

var poolTaobaoAilabAicloudTopHotwordsUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopHotwordsUpdateAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopHotwordsUpdateAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopHotwordsUpdateAPIResponse
func GetTaobaoAilabAicloudTopHotwordsUpdateAPIResponse() *TaobaoAilabAicloudTopHotwordsUpdateAPIResponse {
	return poolTaobaoAilabAicloudTopHotwordsUpdateAPIResponse.Get().(*TaobaoAilabAicloudTopHotwordsUpdateAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopHotwordsUpdateAPIResponse 将 TaobaoAilabAicloudTopHotwordsUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopHotwordsUpdateAPIResponse(v *TaobaoAilabAicloudTopHotwordsUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopHotwordsUpdateAPIResponse.Put(v)
}
