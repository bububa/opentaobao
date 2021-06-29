package aliexpresssumaitong

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
平台固定参数获取 APIResponse
aliexpress.taxation.platform.open.get

Aliexpress开放平台固定参数获取
*/
type AliexpressTaxationPlatformOpenGetAPIResponse struct {
    model.CommonResponse
    AliexpressTaxationPlatformOpenGetResponse
}

type AliexpressTaxationPlatformOpenGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_taxation_platform_open_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ResponseDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
