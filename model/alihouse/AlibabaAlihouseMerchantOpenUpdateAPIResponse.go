package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseMerchantOpenUpdateAPIResponse 非融合店进件升级成融合店 API返回值
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
type AlibabaAlihouseMerchantOpenUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseMerchantOpenUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseMerchantOpenUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseMerchantOpenUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseMerchantOpenUpdateAPIResponseModel is 非融合店进件升级成融合店 成功返回结果
type AlibabaAlihouseMerchantOpenUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_merchant_open_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseMerchantOpenUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseMerchantOpenUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseMerchantOpenUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMerchantOpenUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseMerchantOpenUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseMerchantOpenUpdateAPIResponse
func GetAlibabaAlihouseMerchantOpenUpdateAPIResponse() *AlibabaAlihouseMerchantOpenUpdateAPIResponse {
	return poolAlibabaAlihouseMerchantOpenUpdateAPIResponse.Get().(*AlibabaAlihouseMerchantOpenUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseMerchantOpenUpdateAPIResponse 将 AlibabaAlihouseMerchantOpenUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseMerchantOpenUpdateAPIResponse(v *AlibabaAlihouseMerchantOpenUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseMerchantOpenUpdateAPIResponse.Put(v)
}
