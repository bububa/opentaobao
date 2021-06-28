package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品上架 APIResponse
alibaba.item.operate.upshelf

商品上架
*/
type AlibabaItemOperateUpshelfAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaItemOperateUpshelfResponse `json:"alibaba_item_operate_upshelf_response,omitempty"` 
    AlibabaItemOperateUpshelfResponse
}

/* model for simplify = false
type AlibabaItemOperateUpshelfResponse struct {

    // 商品上架是否成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaItemOperateUpshelfResponse struct {

    // 商品上架是否成功
    Result   string `json:"result,omitempty"`

}
