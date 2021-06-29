package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品发布规则获取接口 APIResponse
tmall.product.add.schema.get

获取用户发布产品的规则
*/
type TmallProductAddSchemaGetAPIResponse struct {
    model.CommonResponse
    TmallProductAddSchemaGetResponse
}

type TmallProductAddSchemaGetResponse struct {
    XMLName xml.Name `xml:"tmall_product_add_schema_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回发布产品的规则文档
    
    AddProductRule   string `json:"add_product_rule,omitempty" xml:"add_product_rule,omitempty"`

    
}
