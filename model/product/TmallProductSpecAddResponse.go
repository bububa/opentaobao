package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加产品规格 APIResponse
tmall.product.spec.add

增加产品规格
*/
type TmallProductSpecAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_product_spec_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品规格对象
    
    ProductSpec   *ProductSpec `json:"product_spec,omitempty" xml:"