package btrip

import (
	"encoding/xml"

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

// AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel is 用车结算记账查询接口 成功返回结果
type AlitripBtripCorpopCarBillsettlementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_car_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
