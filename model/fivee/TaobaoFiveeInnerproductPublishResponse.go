package fivee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国产商品发布 APIResponse
taobao.fivee.innerproduct.publish

资质共享平台国产商品发布
*/
type TaobaoFiveeInnerproductPublishAPIResponse struct {
    model.CommonResponse
    TaobaoFiveeInnerproductPublishResponse
}

type TaobaoFiveeInnerproductPublishResponse struct {
    XMLName xml.Name `xml:"fivee_innerproduct_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // code
    
    CodeT   int64 `json:"code_t,omitempty" xml:"code_t,omitempty"`

    
    // 返回素材id
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 是否成功
    
    SuccessT   bool `json:"success_t,omitempty" xml:"success_t,omitempty"`

    
}
