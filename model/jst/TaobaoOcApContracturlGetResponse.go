package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按用户获取支付宝代扣协议链接地址 APIResponse
taobao.oc.ap.contracturl.get

按用户获取支付宝代扣协议链接地址
*/
type TaobaoOcApContracturlGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOcApContracturlGetResponse `json:"taobao_oc_ap_contracturl_get_response,omitempty"`
}

type TaobaoOcApContracturlGetResponse struct {

    // 判断操作是否执行成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 代扣url地址
    Url   string `json:"url,omitempty"`

    // 错误描述信息
    ErrorDescription   string `json:"error_description,omitempty"`

}
