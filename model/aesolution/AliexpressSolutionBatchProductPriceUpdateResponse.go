package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.price.update API返回值 
aliexpress.solution.batch.product.price.update

batch product price update operation for oversea sellers
*/
type AliexpressSolutionBatchProductPriceUpdateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionBatchProductPriceUpdateResponse
}

// aliexpress.solution.batch.product.price.update 成功返回结果
type AliexpressSolutionBatchProductPriceUpdateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_batch_product_price_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // When success equals false, indicating the error code
    UpdateErrorCode   string `json:"update_error_code,omitempty" xml:"update_error_code,omitempty"`
    // When success equals false, indicating the error message
    UpdateErrorMessage   string `json:"update_error_message,omitempty" xml:"update_error_message,omitempty"`
    // Indicates the update result is successful or not. Only all the products in mutiple_product_update_list have been updated successfully will make the success to be true, otherwise false.
    UpdateSuccess   bool `json:"update_success,omitempty" xml:"update_success,omitempty"`
    // update failed list
    UpdateFailedList   []SynchronizeProductResponseDTO `json:"update_failed_list,omitempty" xml:"update_failed_list>synchronize_product_response_dto,omitempty"`
    // update successful list
    UpdateSuccessfulList   []SynchronizeProductResponseDTO `json:"update_successful_list,omitempty" xml:"update_successful_list>synchronize_product_response_dto,omitempty"`
}
