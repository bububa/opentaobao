package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpophotelbillsettlementqueryAPIResponse 酒店结算记账查询接口 API返回值
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
type AlitripbtripcorpophotelbillsettlementqueryAPIResponse struct {
	model.CommonResponse
	AlitripbtripcorpophotelbillsettlementqueryAPIResponseModel
}

// AlitripbtripcorpophotelbillsettlementqueryAPIResponseModel is 酒店结算记账查询接口 成功返回结果
type AlitripbtripcorpophotelbillsettlementqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_hotel_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
