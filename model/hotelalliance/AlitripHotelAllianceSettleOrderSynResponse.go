package hotelalliance

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菲住联盟分账成功订单同步 APIResponse
alitrip.hotel.alliance.settle.order.syn

用于菲住联盟分账成功订单同步
*/
type AlitripHotelAllianceSettleOrderSynAPIResponse struct {
    model.CommonResponse
    AlitripHotelAllianceSettleOrderSynResponse
}

type AlitripHotelAllianceSettleOrderSynResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_alliance_settle_order_syn_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的结果
    
    HmsTopResultSet   *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`

    
}
