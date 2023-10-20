package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionBatchProductPriceUpdateAPIResponse aliexpress.solution.batch.product.price.update API返回值
// aliexpress.solution.batch.product.price.update
//
// batch product price update operation for oversea sellers
type AliexpressSolutionBatchProductPriceUpdateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionBatchProductPriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionBatchProductPriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionBatchProductPriceUpdateAPIResponseModel).Reset()
}

// AliexpressSolutionBatchProductPriceUpdateAPIResponseModel is aliexpress.solution.batch.product.price.update 成功返回结果
type AliexpressSolutionBatchProductPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_batch_product_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// update failed list
	UpdateFailedList []SynchronizeProductResponseDto `json:"update_failed_list,omitempty" xml:"update_failed_list>synchronize_product_response_dto,omitempty"`
	// update successful list
	UpdateSuccessfulList []SynchronizeProductResponseDto `json:"update_successful_list,omitempty" xml:"update_successful_list>synchronize_product_response_dto,omitempty"`
	// When success equals false, indicating the error code
	UpdateErrorCode string `json:"update_error_code,omitempty" xml:"update_error_code,omitempty"`
	// When success equals false, indicating the error message
	UpdateErrorMessage string `json:"update_error_message,omitempty" xml:"update_error_message,omitempty"`
	// Indicates the update result is successful or not. Only all the products in mutiple_product_update_list have been updated successfully will make the success to be true, otherwise false.
	UpdateSuccess bool `json:"update_success,omitempty" xml:"update_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionBatchProductPriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateFailedList = m.UpdateFailedList[:0]
	m.UpdateSuccessfulList = m.UpdateSuccessfulList[:0]
	m.UpdateErrorCode = ""
	m.UpdateErrorMessage = ""
	m.UpdateSuccess = false
}

var poolAliexpressSolutionBatchProductPriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionBatchProductPriceUpdateAPIResponse)
	},
}

// GetAliexpressSolutionBatchProductPriceUpdateAPIResponse 从 sync.Pool 获取 AliexpressSolutionBatchProductPriceUpdateAPIResponse
func GetAliexpressSolutionBatchProductPriceUpdateAPIResponse() *AliexpressSolutionBatchProductPriceUpdateAPIResponse {
	return poolAliexpressSolutionBatchProductPriceUpdateAPIResponse.Get().(*AliexpressSolutionBatchProductPriceUpdateAPIResponse)
}

// ReleaseAliexpressSolutionBatchProductPriceUpdateAPIResponse 将 AliexpressSolutionBatchProductPriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionBatchProductPriceUpdateAPIResponse(v *AliexpressSolutionBatchProductPriceUpdateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionBatchProductPriceUpdateAPIResponse.Put(v)
}
