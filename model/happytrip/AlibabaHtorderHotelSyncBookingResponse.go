package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店亲橙客栈预订信息同步 APIResponse
alibaba.htorder.hotel.sync.booking

未来酒店亲橙客栈预订信息同步
*/
type AlibabaHtorderHotelSyncBookingAPIResponse struct {
    model.CommonResponse
    AlibabaHtorderHotelSyncBookingResponse
}

type AlibabaHtorderHotelSyncBookingResponse struct {
    XMLName xml.Name `xml:"alibaba_htorder_hotel_sync_booking_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaHtorderHotelSyncBookingResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
