package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopCarBillsettlementQueryAPIResponse 用车结算记账查询接口 API返回值
// alitrip.btrip.corpop.car.billsettlement.query
//
// 用车结算记账查询接口
type AlitripBtripCorpopCarBillsettlementQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopCarBillsettlementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel).Reset()
}

// AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel is 用车结算记账查询接口 成功返回结果
type AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_car_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopCarBillsettlementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopCarBillsettlementQueryAPIResponse)
	},
}

// GetAlitripBtripCorpopCarBillsettlementQueryAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopCarBillsettlementQueryAPIResponse
func GetAlitripBtripCorpopCarBillsettlementQueryAPIResponse() *AlitripBtripCorpopCarBillsettlementQueryAPIResponse {
	return poolAlitripBtripCorpopCarBillsettlementQueryAPIResponse.Get().(*AlitripBtripCorpopCarBillsettlementQueryAPIResponse)
}

// ReleaseAlitripBtripCorpopCarBillsettlementQueryAPIResponse 将 AlitripBtripCorpopCarBillsettlementQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopCarBillsettlementQueryAPIResponse(v *AlitripBtripCorpopCarBillsettlementQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopCarBillsettlementQueryAPIResponse.Put(v)
}
