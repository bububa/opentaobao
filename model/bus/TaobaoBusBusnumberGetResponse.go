package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次查询 APIResponse
taobao.bus.busnumber.get

提供汽车票车次查询服务
*/
type TaobaoBusBusnumberGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusBusnumberGetResponse
}

type TaobaoBusBusnumberGetResponse struct {
    XMLName xml.Name `xml:"bus_busnumber_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoBusBusnumberGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
