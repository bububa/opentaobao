package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenCustomerSaveAPIResponse 保存和更新顾客 API返回值
// alibaba.alsc.crm.open.customer.save
//
// 用来保存顾客，如果已经存在的话，则更新顾客
type AlibabaAlscCrmOpenCustomerSaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenCustomerSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenCustomerSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmOpenCustomerSaveAPIResponseModel).Reset()
}

// AlibabaAlscCrmOpenCustomerSaveAPIResponseModel is 保存和更新顾客 成功返回结果
type AlibabaAlscCrmOpenCustomerSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_customer_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenCustomerSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmOpenCustomerSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmOpenCustomerSaveAPIResponse)
	},
}

// GetAlibabaAlscCrmOpenCustomerSaveAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmOpenCustomerSaveAPIResponse
func GetAlibabaAlscCrmOpenCustomerSaveAPIResponse() *AlibabaAlscCrmOpenCustomerSaveAPIResponse {
	return poolAlibabaAlscCrmOpenCustomerSaveAPIResponse.Get().(*AlibabaAlscCrmOpenCustomerSaveAPIResponse)
}

// ReleaseAlibabaAlscCrmOpenCustomerSaveAPIResponse 将 AlibabaAlscCrmOpenCustomerSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmOpenCustomerSaveAPIResponse(v *AlibabaAlscCrmOpenCustomerSaveAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmOpenCustomerSaveAPIResponse.Put(v)
}
