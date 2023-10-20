package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCustomerAddNewAPIResponse 导购域留资接口 API返回值
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
type AlibabaLsyCrmCustomerAddNewAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmCustomerAddNewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmCustomerAddNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmCustomerAddNewAPIResponseModel).Reset()
}

// AlibabaLsyCrmCustomerAddNewAPIResponseModel is 导购域留资接口 成功返回结果
type AlibabaLsyCrmCustomerAddNewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_customer_add_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLsyCrmCustomerAddNewResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmCustomerAddNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmCustomerAddNewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmCustomerAddNewAPIResponse)
	},
}

// GetAlibabaLsyCrmCustomerAddNewAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmCustomerAddNewAPIResponse
func GetAlibabaLsyCrmCustomerAddNewAPIResponse() *AlibabaLsyCrmCustomerAddNewAPIResponse {
	return poolAlibabaLsyCrmCustomerAddNewAPIResponse.Get().(*AlibabaLsyCrmCustomerAddNewAPIResponse)
}

// ReleaseAlibabaLsyCrmCustomerAddNewAPIResponse 将 AlibabaLsyCrmCustomerAddNewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmCustomerAddNewAPIResponse(v *AlibabaLsyCrmCustomerAddNewAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmCustomerAddNewAPIResponse.Put(v)
}
