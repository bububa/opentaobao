package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新规则获取接口 APIResponse
tmall.product.update.schema.get

获取用户更新产品的规则
*/
type TmallProductUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    TmallProductUpdateSchemaGetResponse
}

type TmallProductUpdateSchemaGetResponse struct {
    XMLName xml.Name `xml:"tmall_product_update_schema_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 参数产品ID对产品的更新规则
    
    UpdateProductSchema   string `json:"update_product_schema,omitempty" xml:"update_product_schema,omitempty"`

    
}
