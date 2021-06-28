package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）获取商品信息 APIResponse
alibaba.icbu.product.schema.render

获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
*/
type AlibabaIcbuProductSchemaRenderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_schema_render_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品发布规则和对应填写数据
    
    Data   string `json:"data,omitempty" xml:"