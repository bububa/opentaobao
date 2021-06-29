package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝自助修改地址服务开通 APIResponse
taobao.modifyaddress.open

商家自助修改地址服务开通
*/
type TaobaoModifyaddressOpenAPIResponse struct {
    model.CommonResponse
    TaobaoModifyaddressOpenResponse
}

type TaobaoModifyaddressOpenResponse struct {
    XMLName xml.Name `xml:"modifyaddress_open_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
