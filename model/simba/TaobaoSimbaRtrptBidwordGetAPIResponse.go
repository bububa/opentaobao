package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptBidwordGetAPIResponse 获取推广词实时报表数据 API返回值
// taobao.simba.rtrpt.bidword.get
//
// 获取推广词报表数据
type TaobaoSimbaRtrptBidwordGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRtrptBidwordGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptBidwordGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRtrptBidwordGetAPIResponseModel).Reset()
}

// TaobaoSimbaRtrptBidwordGetAPIResponseModel is 获取推广词实时报表数据 成功返回结果
type TaobaoSimbaRtrptBidwordGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rtrpt_bidword_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// bidword result
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptBidwordGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRtrptBidwordGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRtrptBidwordGetAPIResponse)
	},
}

// GetTaobaoSimbaRtrptBidwordGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRtrptBidwordGetAPIResponse
func GetTaobaoSimbaRtrptBidwordGetAPIResponse() *TaobaoSimbaRtrptBidwordGetAPIResponse {
	return poolTaobaoSimbaRtrptBidwordGetAPIResponse.Get().(*TaobaoSimbaRtrptBidwordGetAPIResponse)
}

// ReleaseTaobaoSimbaRtrptBidwordGetAPIResponse 将 TaobaoSimbaRtrptBidwordGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRtrptBidwordGetAPIResponse(v *TaobaoSimbaRtrptBidwordGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRtrptBidwordGetAPIResponse.Put(v)
}
