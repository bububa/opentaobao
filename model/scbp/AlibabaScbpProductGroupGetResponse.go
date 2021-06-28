package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定产品分组的下一层子分组 APIResponse
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组
*/
type AlibabaScbpProductGroupGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_product_group_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 下一层分组列表
    
    RoductGroupList   []TopProductGroupDto `json:"roduct_group_list,omitempty" xml:"