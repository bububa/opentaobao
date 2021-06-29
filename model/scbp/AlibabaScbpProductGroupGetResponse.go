package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定产品分组的下一层子分组 API返回值 
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组
*/
type AlibabaScbpProductGroupGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpProductGroupGetResponse
}

// 查询指定产品分组的下一层子分组 成功返回结果
type AlibabaScbpProductGroupGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_product_group_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 下一层分组列表
    RoductGroupList   []TopProductGroupDTO `json:"roduct_group_list,omitempty" xml:"roduct_group_list>top_product_group_dto,omitempty"`
}
