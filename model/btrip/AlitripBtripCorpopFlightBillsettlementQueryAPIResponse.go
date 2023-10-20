package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopFlightBillsettlementQueryAPIResponse 机票结算记账查询接口 API返回值
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
type AlitripBtripCorpopFlightBillsettlementQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopFlightBillsettlementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopFlightBillsettlementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopFlightBillsettlementQueryAPIResponseModel).Reset()
}

// AlitripBtripCorpopFlightBillsettlementQueryAPIResponseModel is 机票结算记账查询接口 成功返回结果
type AlitripBtripCorpopFlightBillsettlementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_flight_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopFlightBillsettlementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopFlightBillsettlementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopFlightBillsettlementQueryAPIResponse)
	},
}

// GetAlitripBtripCorpopFlightBillsettlementQueryAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopFlightBillsettlementQueryAPIResponse
func GetAlitripBtripCorpopFlightBillsettlementQueryAPIResponse() *AlitripBtripCorpopFlightBillsettlementQueryAPIResponse {
	return poolAlitripBtripCorpopFlightBillsettlementQueryAPIResponse.Get().(*AlitripBtripCorpopFlightBillsettlementQueryAPIResponse)
}

// ReleaseAlitripBtripCorpopFlightBillsettlementQueryAPIResponse 将 AlitripBtripCorpopFlightBillsettlementQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopFlightBillsettlementQueryAPIResponse(v *AlitripBtripCorpopFlightBillsettlementQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopFlightBillsettlementQueryAPIResponse.Put(v)
}
