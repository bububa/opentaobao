package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票余票接口 APIResponse
taobao.bus.seatprice.get

提供给商家，查询汽车票班次余票
*/
type TaobaoBusSeatpriceGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusSeatpriceGetResponse
}

type TaobaoBusSeatpriceGetResponse struct {
    XMLName xml.Name `xml:"bus_seatprice_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoBusSeatpriceGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
