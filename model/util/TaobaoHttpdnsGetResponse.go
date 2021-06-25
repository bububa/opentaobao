package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
TOPDNS配置 APIResponse
taobao.httpdns.get

获取TOP DNS配置
*/
type TaobaoHttpdnsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoHttpdnsGetResponse `json:"taobao_httpdns_get_response,omitempty"`
}

type TaobaoHttpdnsGetResponse struct {

    // HTTP DNS配置信息
    Result   string `json:"result,omitempty"`

}
