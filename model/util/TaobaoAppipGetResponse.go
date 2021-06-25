package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取ISV发起请求服务器IP APIResponse
taobao.appip.get

获取ISV发起请求服务器IP
*/
type TaobaoAppipGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAppipGetResponse `json:"taobao_appip_get_response,omitempty"`
}

type TaobaoAppipGetResponse struct {

    // ISV发起请求服务器IP
    Ip   string `json:"ip,omitempty"`

}
