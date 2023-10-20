package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLafiteSellerBenefitListAPIResponse 商家自运营权益列表 API返回值
// alibaba.lafite.seller.benefit.list
//
// 小程序isv可使用该接口获取权益列表
type AlibabaLafiteSellerBenefitListAPIResponse struct {
	model.CommonResponse
	AlibabaLafiteSellerBenefitListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLafiteSellerBenefitListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLafiteSellerBenefitListAPIResponseModel).Reset()
}

// AlibabaLafiteSellerBenefitListAPIResponseModel is 商家自运营权益列表 成功返回结果
type AlibabaLafiteSellerBenefitListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lafite_seller_benefit_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLafiteSellerBenefitListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLafiteSellerBenefitListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLafiteSellerBenefitListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLafiteSellerBenefitListAPIResponse)
	},
}

// GetAlibabaLafiteSellerBenefitListAPIResponse 从 sync.Pool 获取 AlibabaLafiteSellerBenefitListAPIResponse
func GetAlibabaLafiteSellerBenefitListAPIResponse() *AlibabaLafiteSellerBenefitListAPIResponse {
	return poolAlibabaLafiteSellerBenefitListAPIResponse.Get().(*AlibabaLafiteSellerBenefitListAPIResponse)
}

// ReleaseAlibabaLafiteSellerBenefitListAPIResponse 将 AlibabaLafiteSellerBenefitListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLafiteSellerBenefitListAPIResponse(v *AlibabaLafiteSellerBenefitListAPIResponse) {
	v.Reset()
	poolAlibabaLafiteSellerBenefitListAPIResponse.Put(v)
}
