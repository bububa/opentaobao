package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptCustGetAPIResponse 获取账户实时报表数据 API返回值
// taobao.simba.rtrpt.cust.get
//
// 获取账户实时报表数据
type TaobaoSimbaRtrptCustGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRtrptCustGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptCustGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRtrptCustGetAPIResponseModel).Reset()
}

// TaobaoSimbaRtrptCustGetAPIResponseModel is 获取账户实时报表数据 成功返回结果
type TaobaoSimbaRtrptCustGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rtrpt_cust_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptCustGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoSimbaRtrptCustGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRtrptCustGetAPIResponse)
	},
}

// GetTaobaoSimbaRtrptCustGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRtrptCustGetAPIResponse
func GetTaobaoSimbaRtrptCustGetAPIResponse() *TaobaoSimbaRtrptCustGetAPIResponse {
	return poolTaobaoSimbaRtrptCustGetAPIResponse.Get().(*TaobaoSimbaRtrptCustGetAPIResponse)
}

// ReleaseTaobaoSimbaRtrptCustGetAPIResponse 将 TaobaoSimbaRtrptCustGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRtrptCustGetAPIResponse(v *TaobaoSimbaRtrptCustGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRtrptCustGetAPIResponse.Put(v)
}
