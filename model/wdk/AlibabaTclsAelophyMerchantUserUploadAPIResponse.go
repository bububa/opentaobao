package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantUserUploadAPIResponse 商家会员数据上传 API返回值
// alibaba.tcls.aelophy.merchant.user.upload
//
// 商家会员数据上传
type AlibabaTclsAelophyMerchantUserUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantUserUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantUserUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantUserUploadAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantUserUploadAPIResponseModel is 商家会员数据上传 成功返回结果
type AlibabaTclsAelophyMerchantUserUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_user_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	ApiResult *AlibabaTclsAelophyMerchantUserUploadApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantUserUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantUserUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantUserUploadAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantUserUploadAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantUserUploadAPIResponse
func GetAlibabaTclsAelophyMerchantUserUploadAPIResponse() *AlibabaTclsAelophyMerchantUserUploadAPIResponse {
	return poolAlibabaTclsAelophyMerchantUserUploadAPIResponse.Get().(*AlibabaTclsAelophyMerchantUserUploadAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantUserUploadAPIResponse 将 AlibabaTclsAelophyMerchantUserUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantUserUploadAPIResponse(v *AlibabaTclsAelophyMerchantUserUploadAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantUserUploadAPIResponse.Put(v)
}
