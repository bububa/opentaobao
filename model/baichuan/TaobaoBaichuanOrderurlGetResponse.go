package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川订单详情 APIResponse
taobao.baichuan.orderurl.get

百川订单详情
*/
type TaobaoBaichuanOrderurlGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOrderurlGetResponse `json:"baichuan_orderurl_get_response,omitempty"` 
    TaobaoBaichuanOrderurlGetResponse
}

/* model for simplify = false
type TaobaoBaichuanOrderurlGetResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOrderurlGetResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
