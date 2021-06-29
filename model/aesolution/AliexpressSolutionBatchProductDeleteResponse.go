package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.delete APIResponse
aliexpress.solution.batch.product.delete

Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
*/
type AliexpressSolutionBatchProductDeleteAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionBatchProductDeleteResponse
}

type AliexpressSolutionBatchProductDeleteResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_batch_product_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
