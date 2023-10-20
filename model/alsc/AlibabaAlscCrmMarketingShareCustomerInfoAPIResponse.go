package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse 查询分享营销客户领券信息 API返回值
// alibaba.alsc.crm.marketing.share.customer.info
//
// 查询分享营销活动的客户领券信息
type AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmMarketingShareCustomerInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmMarketingShareCustomerInfoAPIResponseModel).Reset()
}

// AlibabaAlscCrmMarketingShareCustomerInfoAPIResponseModel is 查询分享营销客户领券信息 成功返回结果
type AlibabaAlscCrmMarketingShareCustomerInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_marketing_share_customer_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmMarketingShareCustomerInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse)
	},
}

// GetAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse
func GetAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse() *AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse {
	return poolAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse.Get().(*AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse)
}

// ReleaseAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse 将 AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse(v *AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmMarketingShareCustomerInfoAPIResponse.Put(v)
}
