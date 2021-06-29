package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.inventory.update APIResponse
aliexpress.solution.batch.product.inventory.update

batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product.
*/
type AliexpressSolutionBatchProductInventoryUpdateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionBatchProductInventoryUpdateResponse
}

type AliexpressSolutionBatchProductInventoryUpdateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_batch_product_inventory_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // When success equals false, indicating the error code
    
    UpdateErrorCode   string `json:"update_error_code,omitempty" xml:"update_error_code,omitempty"`

    
    // When success equals false,  indicating the error message
    
    UpdateErrorMessage   string `json:"update_error_message,omitempty" xml:"update_error_message,omitempty"`

    
    // Indicates the update result is successful or not. Only all the products in mutiple_product_update_list have been updated successfully will make the success to be true, otherwise false.
    
    UpdateSuccess   bool `json:"update_success,omitempty" xml:"update_success,omitempty"`

    
    // update failed list
    
    UpdateFailedList   []SynchronizeProductResponseDto `json:"update_failed_list,omitempty" xml:"update_failed_list>synchronize_product_response_dto,omitempty"`
    
    
    // update succesful list
    
    UpdateSuccessfulList   []SynchronizeProductResponseDto `json:"update_successful_list,omitempty" xml:"update_successful_list>synchronize_product_response_dto,omitempty"`
    
    
}
