package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分组信息获取 APIResponse
alibaba.icbu.product.group.get

分组信息获取
*/
type AlibabaIcbuProductGroupGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductGroupGetResponse `json:"alibaba_icbu_product_group_get_response,omitempty"` 
    AlibabaIcbuProductGroupGetResponse
}

/* model for simplify = false
type AlibabaIcbuProductGroupGetResponse struct {

    // 分组信息
    
    ProductGroup  *struct {
        ProductGroup  *ProductGroup `json:"product_group,omitempty"`
    } `json:"product_group,omitempty"`
    

}
*/

type AlibabaIcbuProductGroupGetResponse struct {

    // 分组信息
    ProductGroup   *ProductGroup `json:"product_group,omitempty"`

}
