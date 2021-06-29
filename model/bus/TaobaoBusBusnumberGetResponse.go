package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次查询 API返回值 
taobao.bus.busnumber.get

提供汽车票车次查询服务
*/
type TaobaoBusBusnumberGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusBusnumberGetResponse
}

// 汽车票车次查询 成功返回结果
type TaobaoBusBusnumberGetResponse struct {
    XMLName xml.Name `xml:"bus_busnumber_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoBusBusnumberGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
