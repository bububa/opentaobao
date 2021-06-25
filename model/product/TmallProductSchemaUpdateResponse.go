package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 APIResponse
tmall.product.schema.update

产品更新接口
*/
type TmallProductSchemaUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallProductSchemaUpdateResponse `json:"tmall_product_schema_update_response,omitempty"`
}

type TmallProductSchemaUpdateResponse struct {

    // 产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间
    UpdateProductResult   string `json:"update_product_result,omitempty"`

}
