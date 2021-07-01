package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票余票接口 API返回值 
taobao.bus.seatprice.get

提供给商家，查询汽车票班次余票
*/
type TaobaoBusSeatpriceGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusSeatpriceGetAPIResponseModel
}

// 汽车票余票接口 成功返回结果
type TaobaoBusSeatpriceGetAPIResponseModel struct {
    XMLName xml.Name `xml:"bus_seatprice_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoBusSeatpriceGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
