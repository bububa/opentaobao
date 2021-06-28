package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川订单详情 APIResponse
taobao.baichuan.orderurl.get

百川订单详情
*/
type TaobaoBaichuanOrderurlGetAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOrderurlGetResponse
}

type TaobaoBaichuanOrderurlGetResponse struct {
    XMLName xml.Name `xml:"baichuan_orderurl_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
