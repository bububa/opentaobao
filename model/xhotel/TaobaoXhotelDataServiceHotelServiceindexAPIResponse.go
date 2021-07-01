package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店服务指数 API返回值 
taobao.xhotel.data.service.hotel.serviceindex

酒店服务指数
*/
type TaobaoXhotelDataServiceHotelServiceindexAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel
}

// 酒店服务指数 成功返回结果
type TaobaoXhotelDataServiceHotelServiceindexAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_data_service_hotel_serviceindex_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoXhotelDataServiceHotelServiceindexResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
