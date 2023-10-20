package hotelalliance

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelalliancesettleordersynAPIResponse 菲住联盟分账成功订单同步 API返回值
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
type AlitriphotelalliancesettleordersynAPIResponse struct {
	model.CommonResponse
	AlitriphotelalliancesettleordersynAPIResponseModel
}

// AlitriphotelalliancesettleordersynAPIResponseModel is 菲住联盟分账成功订单同步 成功返回结果
type AlitriphotelalliancesettleordersynAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_alliance_settle_order_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的结果
	HmsTopResultSet *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`
}
