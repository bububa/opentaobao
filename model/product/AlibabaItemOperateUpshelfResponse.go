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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_item_operate_upshelf_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品上架是否成功
    
    Result   string `json:"result,omitempty" xml:"