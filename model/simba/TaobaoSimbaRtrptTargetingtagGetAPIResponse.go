package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptTargetingtagGetAPIResponse 搜索人群实时报表 API返回值
// taobao.simba.rtrpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaoSimbaRtrptTargetingtagGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRtrptTargetingtagGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptTargetingtagGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRtrptTargetingtagGetAPIResponseModel).Reset()
}

// TaobaoSimbaRtrptTargetingtagGetAPIResponseModel is 搜索人群实时报表 成功返回结果
type TaobaoSimbaRtrptTargetingtagGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rtrpt_targetingtag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 111
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptTargetingtagGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRtrptTargetingtagGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRtrptTargetingtagGetAPIResponse)
	},
}

// GetTaobaoSimbaRtrptTargetingtagGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRtrptTargetingtagGetAPIResponse
func GetTaobaoSimbaRtrptTargetingtagGetAPIResponse() *TaobaoSimbaRtrptTargetingtagGetAPIResponse {
	return poolTaobaoSimbaRtrptTargetingtagGetAPIResponse.Get().(*TaobaoSimbaRtrptTargetingtagGetAPIResponse)
}

// ReleaseTaobaoSimbaRtrptTargetingtagGetAPIResponse 将 TaobaoSimbaRtrptTargetingtagGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRtrptTargetingtagGetAPIResponse(v *TaobaoSimbaRtrptTargetingtagGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRtrptTargetingtagGetAPIResponse.Put(v)
}
