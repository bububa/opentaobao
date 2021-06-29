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
    TaobaoOcApContracturlGetResponse
}

type TaobaoOcApContracturlGetResponse struct {
    XMLName xml.Name `xml:"oc_ap_contracturl_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 判断操作是否执行成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 代扣url地址
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`

    
    // 错误描述信息
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
}
