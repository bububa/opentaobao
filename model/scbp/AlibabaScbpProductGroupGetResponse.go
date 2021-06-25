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
    Response *AlibabaScbpProductGroupGetResponse `json:"alibaba_scbp_product_group_get_response,omitempty"`
}

type AlibabaScbpProductGroupGetResponse struct {

    // 下一层分组列表
    RoductGroupList   []TopProductGroupDto `json:"roduct_group_list,omitempty"`

}
