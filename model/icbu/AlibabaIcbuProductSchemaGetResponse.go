package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布schema接口 APIResponse
alibaba.icbu.product.schema.get

获取ICBU商品发布的页面规则和填写字段，适用于新发商品
*/
type AlibabaIcbuProductSchemaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_schema_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品发布规则
    
    Data   string `json:"data,omitempty" xml:"