package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新接口 APIResponse
tmall.product.schema.update

产品更新接口
*/
type TmallProductSchemaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_product_schema_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间
    
    UpdateProductResult   string `json:"update_product_result,omitempty" xml:"