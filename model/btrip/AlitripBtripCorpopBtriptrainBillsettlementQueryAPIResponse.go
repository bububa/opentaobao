package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponse 商旅火车票结算记账查询接口 API返回值
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
type AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponse struct {
	model.CommonResponse
	AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponseModel
}

// AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponseModel is 商旅火车票结算记账查询接口 成功返回结果
type AlitripbtripcorpopbtriptrainbillsettlementqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_btriptrain_billsettlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
