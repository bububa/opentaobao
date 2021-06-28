package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）商品发布新接口 APIResponse
alibaba.icbu.product.schema.add

提供发布ICBU商品的入口
*/
type AlibabaIcbuProductSchemaAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_schema_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品明文id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"