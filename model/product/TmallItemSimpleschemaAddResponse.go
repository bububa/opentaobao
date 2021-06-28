package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化发布商品 APIResponse
tmall.item.simpleschema.add

天猫简化版schema发布商品。
*/
type TmallItemSimpleschemaAddAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemSimpleschemaAddResponse `json:"tmall_item_simpleschema_add_response,omitempty"` 
    TmallItemSimpleschemaAddResponse
}

/* model for simplify = false
type TmallItemSimpleschemaAddResponse struct {

    // 发布成功后返回商品ID
    
    Result   string `json:"result,omitempty"`
    

    // 商品最后发布时间。
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

}
*/

type TmallItemSimpleschemaAddResponse struct {

    // 发布成功后返回商品ID
    Result   string `json:"result,omitempty"`

    // 商品最后发布时间。
    GmtModified   string `json:"gmt_modified,omitempty"`

}
