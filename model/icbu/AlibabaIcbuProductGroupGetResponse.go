package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分组信息获取 APIResponse
alibaba.icbu.product.group.get

分组信息获取
*/
type AlibabaIcbuProductGroupGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_group_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分组信息
    
    ProductGroup   *ProductGroup `json:"product_group,omitempty" xml:"