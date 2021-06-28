package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 APIResponse
alibaba.item.operate.downshelf

商品下架
*/
type AlibabaItemOperateDownshelfAPIResponse struct {
    model.CommonResponse
    AlibabaItemOperateDownshelfResponse
}

type AlibabaItemOperateDownshelfResponse struct {
    XMLName xml.Name `xml:"alibaba_item_operate_downshelf_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品下架是否成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
