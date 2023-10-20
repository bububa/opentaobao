package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantOrderUploadAPIResponse 商家订单数据上传 API返回值
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
type AlibabaTclsAelophyMerchantOrderUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantOrderUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantOrderUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantOrderUploadAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantOrderUploadAPIResponseModel is 商家订单数据上传 成功返回结果
type AlibabaTclsAelophyMerchantOrderUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabaTclsAelophyMerchantOrderUploadApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantOrderUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantOrderUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantOrderUploadAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantOrderUploadAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantOrderUploadAPIResponse
func GetAlibabaTclsAelophyMerchantOrderUploadAPIResponse() *AlibabaTclsAelophyMerchantOrderUploadAPIResponse {
	return poolAlibabaTclsAelophyMerchantOrderUploadAPIResponse.Get().(*AlibabaTclsAelophyMerchantOrderUploadAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantOrderUploadAPIResponse 将 AlibabaTclsAelophyMerchantOrderUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantOrderUploadAPIResponse(v *AlibabaTclsAelophyMerchantOrderUploadAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantOrderUploadAPIResponse.Put(v)
}
