package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品上架 APIResponse
alibaba.item.operate.upshelf

商品上架
*/
type AlibabaItemOperateUpshelfAPIResponse struct {
    model.CommonResponse
    AlibabaItemOperateUpshelfResponse
}

type AlibabaItemOperateUpshelfResponse struct {
    XMLName xml.Name `xml:"alibaba_item_operate_upshelf_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品上架是否成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
