package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签显示用商品信息写入 APIResponse
taobao.uscesl.iteminfo.put

用于电子价签上显示的商品信息的写入，包含价格及促销信息
*/
type TaobaoUsceslIteminfoPutAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslIteminfoPutResponse
}

type TaobaoUsceslIteminfoPutResponse struct {
    XMLName xml.Name `xml:"uscesl_iteminfo_put_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // target
    
    Target   bool `json:"target,omitempty" xml:"target,omitempty"`

    
    // returnCode
    
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
