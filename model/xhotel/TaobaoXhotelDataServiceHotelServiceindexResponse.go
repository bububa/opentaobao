package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
酒店服务指数 APIResponse
taobao.xhotel.data.service.hotel.serviceindex

酒店服务指数
*/
type TaobaoXhotelDataServiceHotelServiceindexAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelDataServiceHotelServiceindexResponse `json:"xhotel_data_service_hotel_serviceindex_response,omitempty"` 
    TaobaoXhotelDataServiceHotelServiceindexResponse
}

/* model for simplify = false
type TaobaoXhotelDataServiceHotelServiceindexResponse struct {

    // result
    
    Result  *struct {
        TaobaoXhotelDataServiceHotelServiceindexResultSet  *TaobaoXhotelDataServiceHotelServiceindexResultSet `json:"taobao_xhotel_data_service_hotel_serviceindex_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoXhotelDataServiceHotelServiceindexResponse struct {

    // result
    Result   *TaobaoXhotelDataServiceHotelServiceindexResultSet `json:"result,omitempty"`

}
