package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加商品分组 APIResponse
alibaba.icbu.product.group.add

增加商品分组
*/
type AlibabaIcbuProductGroupAddAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductGroupAddResponse
}

type AlibabaIcbuProductGroupAddResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_group_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建的分组信息
    
    ProductGroup   *ProductGroup `json:"product_group,omitempty" xml:"product_group,omitempty"`

    
}
