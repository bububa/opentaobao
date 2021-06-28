package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询指定产品分组的下一层子分组 APIResponse
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组
*/
type AlibabaScbpProductGroupGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpProductGroupGetResponse `json:"alibaba_scbp_product_group_get_response,omitempty"` 
    AlibabaScbpProductGroupGetResponse
}

/* model for simplify = false
type AlibabaScbpProductGroupGetResponse struct {

    // 下一层分组列表
    
    RoductGroupList  struct {
        TopProductGroupDto  []TopProductGroupDto `json:"top_product_group_dto,omitempty"`
    } `json:"roduct_group_list,omitempty"`
    

}
*/

type AlibabaScbpProductGroupGetResponse struct {

    // 下一层分组列表
    RoductGroupList   []TopProductGroupDto `json:"roduct_group_list,omitempty"`

}
