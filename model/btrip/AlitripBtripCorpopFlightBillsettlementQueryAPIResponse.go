package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopflightbillsettlementqueryAPIResponse 机票结算记账查询接口 API返回值
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
type AlitripbtripcorpopflightbillsettlementqueryAPIResponse struct {
	model.CommonResponse
	AlitripbtripcorpopflightbillsettlementqueryAPIResponseModel
}

// AlitripbtripcorpopflightbillsettlementqueryAPIResponseModel is 机票结算记账查询接口 成功返回结果
type AlitripbtripcorpopflightbillsettlementqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_flight_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
