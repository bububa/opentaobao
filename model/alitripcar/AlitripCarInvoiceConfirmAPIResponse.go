package alitripcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripCarInvoiceConfirmAPIResponse 发票确认接口 API返回值
// alitrip.car.invoice.confirm
//
// 飞猪发票回调接口
type AlitripCarInvoiceConfirmAPIResponse struct {
	model.CommonResponse
	AlitripCarInvoiceConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripCarInvoiceConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripCarInvoiceConfirmAPIResponseModel).Reset()
}

// AlitripCarInvoiceConfirmAPIResponseModel is 发票确认接口 成功返回结果
type AlitripCarInvoiceConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_invoice_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
	// 结果对象
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlitripCarInvoiceConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MessageCode = 0
	m.Model = false
}

var poolAlitripCarInvoiceConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripCarInvoiceConfirmAPIResponse)
	},
}

// GetAlitripCarInvoiceConfirmAPIResponse 从 sync.Pool 获取 AlitripCarInvoiceConfirmAPIResponse
func GetAlitripCarInvoiceConfirmAPIResponse() *AlitripCarInvoiceConfirmAPIResponse {
	return poolAlitripCarInvoiceConfirmAPIResponse.Get().(*AlitripCarInvoiceConfirmAPIResponse)
}

// ReleaseAlitripCarInvoiceConfirmAPIResponse 将 AlitripCarInvoiceConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlitripCarInvoiceConfirmAPIResponse(v *AlitripCarInvoiceConfirmAPIResponse) {
	v.Reset()
	poolAlitripCarInvoiceConfirmAPIResponse.Put(v)
}
