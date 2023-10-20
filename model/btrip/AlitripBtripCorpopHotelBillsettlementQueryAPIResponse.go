package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopHotelBillsettlementQueryAPIResponse 酒店结算记账查询接口 API返回值
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
type AlitripBtripCorpopHotelBillsettlementQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopHotelBillsettlementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopHotelBillsettlementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopHotelBillsettlementQueryAPIResponseModel).Reset()
}

// AlitripBtripCorpopHotelBillsettlementQueryAPIResponseModel is 酒店结算记账查询接口 成功返回结果
type AlitripBtripCorpopHotelBillsettlementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_hotel_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopHotelBillsettlementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopHotelBillsettlementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopHotelBillsettlementQueryAPIResponse)
	},
}

// GetAlitripBtripCorpopHotelBillsettlementQueryAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopHotelBillsettlementQueryAPIResponse
func GetAlitripBtripCorpopHotelBillsettlementQueryAPIResponse() *AlitripBtripCorpopHotelBillsettlementQueryAPIResponse {
	return poolAlitripBtripCorpopHotelBillsettlementQueryAPIResponse.Get().(*AlitripBtripCorpopHotelBillsettlementQueryAPIResponse)
}

// ReleaseAlitripBtripCorpopHotelBillsettlementQueryAPIResponse 将 AlitripBtripCorpopHotelBillsettlementQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopHotelBillsettlementQueryAPIResponse(v *AlitripBtripCorpopHotelBillsettlementQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopHotelBillsettlementQueryAPIResponse.Put(v)
}
