package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加商品分组 APIResponse
alibaba.icbu.product.group.add

增加商品分组
*/
type AlibabaIcbuProductGroupAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductGroupAddResponse `json:"alibaba_icbu_product_group_add_response,omitempty"` 
    AlibabaIcbuProductGroupAddResponse
}

/* model for simplify = false
type AlibabaIcbuProductGroupAddResponse struct {

    // 创建的分组信息
    
    ProductGroup  *struct {
        ProductGroup  *ProductGroup `json:"product_group,omitempty"`
    } `json:"product_group,omitempty"`
    

}
*/

type AlibabaIcbuProductGroupAddResponse struct {

    // 创建的分组信息
    ProductGroup   *ProductGroup `json:"product_group,omitempty"`

}
