package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.schema.product.full.update APIResponse
aliexpress.solution.schema.product.full.update

Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
*/
type AliexpressSolutionSchemaProductFullUpdateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionSchemaProductFullUpdateResponse
}

type AliexpressSolutionSchemaProductFullUpdateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_schema_product_full_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Product id that has been updated.
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`

    
}
