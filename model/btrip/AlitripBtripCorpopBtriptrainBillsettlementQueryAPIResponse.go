package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse 商旅火车票结算记账查询接口 API返回值
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
type AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponseModel).Reset()
}

// AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponseModel is 商旅火车票结算记账查询接口 成功返回结果
type AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_btriptrain_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse)
	},
}

// GetAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse
func GetAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse() *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse {
	return poolAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse.Get().(*AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse)
}

// ReleaseAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse 将 AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse(v *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopBtriptrainBillsettlementQueryAPIResponse.Put(v)
}
