package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCustomerAddAPIResponse 私域导购添加活动留资入口 API返回值
// alibaba.lsy.crm.customer.add
//
// 私域导购添加活动留资入口
type AlibabaLsyCrmCustomerAddAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmCustomerAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmCustomerAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmCustomerAddAPIResponseModel).Reset()
}

// AlibabaLsyCrmCustomerAddAPIResponseModel is 私域导购添加活动留资入口 成功返回结果
type AlibabaLsyCrmCustomerAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_customer_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlibabaLsyCrmCustomerAddResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmCustomerAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmCustomerAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmCustomerAddAPIResponse)
	},
}

// GetAlibabaLsyCrmCustomerAddAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmCustomerAddAPIResponse
func GetAlibabaLsyCrmCustomerAddAPIResponse() *AlibabaLsyCrmCustomerAddAPIResponse {
	return poolAlibabaLsyCrmCustomerAddAPIResponse.Get().(*AlibabaLsyCrmCustomerAddAPIResponse)
}

// ReleaseAlibabaLsyCrmCustomerAddAPIResponse 将 AlibabaLsyCrmCustomerAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmCustomerAddAPIResponse(v *AlibabaLsyCrmCustomerAddAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmCustomerAddAPIResponse.Put(v)
}
