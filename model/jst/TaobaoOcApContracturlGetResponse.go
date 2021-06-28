package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按用户获取支付宝代扣协议链接地址 APIResponse
taobao.oc.ap.contracturl.get

按用户获取支付宝代扣协议链接地址
*/
type TaobaoOcApContracturlGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"oc_ap_contracturl_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 判断操作是否执行成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"