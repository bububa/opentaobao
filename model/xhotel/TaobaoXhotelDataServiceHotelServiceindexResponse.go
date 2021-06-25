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
    Response *TaobaoXhotelDataServiceHotelServiceindexResponse `json:"taobao_xhotel_data_service_hotel_serviceindex_response,omitempty"`
}

type TaobaoXhotelDataServiceHotelServiceindexResponse struct {

    // result
    Result   *TaobaoXhotelDataServiceHotelServiceindexResultSet `json:"result,omitempty"`

}
