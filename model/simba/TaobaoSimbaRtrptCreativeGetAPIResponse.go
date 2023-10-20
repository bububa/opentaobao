package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptCreativeGetAPIResponse 获取创意实时报表数据 API返回值
// taobao.simba.rtrpt.creative.get
//
// 获取创意实时报表数据
type TaobaoSimbaRtrptCreativeGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRtrptCreativeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptCreativeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRtrptCreativeGetAPIResponseModel).Reset()
}

// TaobaoSimbaRtrptCreativeGetAPIResponseModel is 获取创意实时报表数据 成功返回结果
type TaobaoSimbaRtrptCreativeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rtrpt_creative_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 111
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptCreativeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRtrptCreativeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRtrptCreativeGetAPIResponse)
	},
}

// GetTaobaoSimbaRtrptCreativeGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRtrptCreativeGetAPIResponse
func GetTaobaoSimbaRtrptCreativeGetAPIResponse() *TaobaoSimbaRtrptCreativeGetAPIResponse {
	return poolTaobaoSimbaRtrptCreativeGetAPIResponse.Get().(*TaobaoSimbaRtrptCreativeGetAPIResponse)
}

// ReleaseTaobaoSimbaRtrptCreativeGetAPIResponse 将 TaobaoSimbaRtrptCreativeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRtrptCreativeGetAPIResponse(v *TaobaoSimbaRtrptCreativeGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRtrptCreativeGetAPIResponse.Put(v)
}
