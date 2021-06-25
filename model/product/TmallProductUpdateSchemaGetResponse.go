package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品更新规则获取接口 APIResponse
tmall.product.update.schema.get

获取用户更新产品的规则
*/
type TmallProductUpdateSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *TmallProductUpdateSchemaGetResponse `json:"tmall_product_update_schema_get_response,omitempty"`
}

type TmallProductUpdateSchemaGetResponse struct {

    // 参数产品ID对产品的更新规则
    UpdateProductSchema   string `json:"update_product_schema,omitempty"`

}
