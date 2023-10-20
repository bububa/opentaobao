package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCustomerAppointAPIResponse ISV直播间预约 API返回值
// alibaba.lsy.crm.customer.appoint
//
// ISV直播间预约
type AlibabaLsyCrmCustomerAppointAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmCustomerAppointAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmCustomerAppointAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmCustomerAppointAPIResponseModel).Reset()
}

// AlibabaLsyCrmCustomerAppointAPIResponseModel is ISV直播间预约 成功返回结果
type AlibabaLsyCrmCustomerAppointAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_customer_appoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// asdasd
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmCustomerAppointAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmCustomerAppointAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmCustomerAppointAPIResponse)
	},
}

// GetAlibabaLsyCrmCustomerAppointAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmCustomerAppointAPIResponse
func GetAlibabaLsyCrmCustomerAppointAPIResponse() *AlibabaLsyCrmCustomerAppointAPIResponse {
	return poolAlibabaLsyCrmCustomerAppointAPIResponse.Get().(*AlibabaLsyCrmCustomerAppointAPIResponse)
}

// ReleaseAlibabaLsyCrmCustomerAppointAPIResponse 将 AlibabaLsyCrmCustomerAppointAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmCustomerAppointAPIResponse(v *AlibabaLsyCrmCustomerAppointAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmCustomerAppointAPIResponse.Put(v)
}
