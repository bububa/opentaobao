package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布草稿 APIResponse
alibaba.icbu.product.schema.add.draft

提供发布ICBU商品草稿的入口
*/
type AlibabaIcbuProductSchemaAddDraftAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_schema_add_draft_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品草稿明文id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"