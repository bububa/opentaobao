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
    AlibabaIcbuProductGroupGetResponse
}

type AlibabaIcbuProductGroupGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_group_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分组信息
    
    ProductGroup   *ProductGroup `json:"product_group,omitempty" xml:"product_group,omitempty"`

    
}
