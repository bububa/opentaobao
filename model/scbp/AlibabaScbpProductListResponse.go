package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询P4P产品 APIResponse
alibaba.scbp.product.list

查询P4P产品
*/
type AlibabaScbpProductListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpProductListResponse `json:"alibaba_scbp_product_list_response,omitempty"` 
    AlibabaScbpProductListResponse
}

/* model for simplify = false
type AlibabaScbpProductListResponse struct {

    // 产品列表
    
    ProductList  struct {
        TopProductDto  []TopProductDto `json:"top_product_dto,omitempty"`
    } `json:"product_list,omitempty"`
    

    // 总数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

}
*/

type AlibabaScbpProductListResponse struct {

    // 产品列表
    ProductList   []TopProductDto `json:"product_list,omitempty"`

    // 总数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

}
