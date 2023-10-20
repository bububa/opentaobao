package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopHotwordsGetAPIResponse 获取热词 API返回值
// taobao.ailab.aicloud.top.hotwords.get
//
// 获取ASR热词
type TaobaoAilabAicloudTopHotwordsGetAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopHotwordsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopHotwordsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopHotwordsGetAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopHotwordsGetAPIResponseModel is 获取热词 成功返回结果
type TaobaoAilabAicloudTopHotwordsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_hotwords_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// baseresult
	Baseresult *BaseResult `json:"baseresult,omitempty" xml:"baseresult,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopHotwordsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Baseresult = nil
}

var poolTaobaoAilabAicloudTopHotwordsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopHotwordsGetAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopHotwordsGetAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopHotwordsGetAPIResponse
func GetTaobaoAilabAicloudTopHotwordsGetAPIResponse() *TaobaoAilabAicloudTopHotwordsGetAPIResponse {
	return poolTaobaoAilabAicloudTopHotwordsGetAPIResponse.Get().(*TaobaoAilabAicloudTopHotwordsGetAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopHotwordsGetAPIResponse 将 TaobaoAilabAicloudTopHotwordsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopHotwordsGetAPIResponse(v *TaobaoAilabAicloudTopHotwordsGetAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopHotwordsGetAPIResponse.Put(v)
}
