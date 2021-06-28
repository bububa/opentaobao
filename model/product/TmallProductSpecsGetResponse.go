package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品的规格信息 APIResponse
tmall.product.specs.get

按product_id或品牌下载产品规格，返回一组的产品规格信息。
*/
type TmallProductSpecsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_product_specs_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回一组产品规格信息。
    
    ProductSpecs   []ProductSpec `json:"product_specs,omitempty" xml:"