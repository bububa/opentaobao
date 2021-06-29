package fivee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
进口商品发布 APIResponse
taobao.fivee.importproduct.publish

直营业务商家入住发布商品时，上传商品及商家证照信息
*/
type TaobaoFiveeImportproductPublishAPIResponse struct {
    model.CommonResponse
    TaobaoFiveeImportproductPublishResponse
}

type TaobaoFiveeImportproductPublishResponse struct {
    XMLName xml.Name `xml:"fivee_importproduct_publish_response"`
    
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
