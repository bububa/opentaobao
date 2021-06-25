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
    Response *TaobaoBaichuanOrderurlGetResponse `json:"taobao_baichuan_orderurl_get_response,omitempty"`
}

type TaobaoBaichuanOrderurlGetResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
