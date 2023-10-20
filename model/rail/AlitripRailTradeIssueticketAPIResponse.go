package rail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriprailtradeissueticketAPIResponse 德铁出票成功接口 API返回值
// alitrip.rail.trade.issueticket
//
// 出票成功回调接口
type AlitriprailtradeissueticketAPIResponse struct {
	model.CommonResponse
	AlitriprailtradeissueticketAPIResponseModel
}

// AlitriprailtradeissueticketAPIResponseModel is 德铁出票成功接口 成功返回结果
type AlitriprailtradeissueticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_trade_issueticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *AlitriprailtradeissueticketResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
