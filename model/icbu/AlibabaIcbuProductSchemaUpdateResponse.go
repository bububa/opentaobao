package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）商品发布增量更新接口 APIResponse
alibaba.icbu.product.schema.update

商品更新接口，方式为增量更新，只更新传入的字段
*/
type AlibabaIcbuProductSchemaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_schema_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品明文id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"