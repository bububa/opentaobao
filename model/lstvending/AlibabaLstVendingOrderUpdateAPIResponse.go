package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingOrderUpdateAPIResponse 自动售货机订单物流信息回传 API返回值
// alibaba.lst.vending.order.update
//
// 零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
type AlibabaLstVendingOrderUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingOrderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingOrderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingOrderUpdateAPIResponseModel).Reset()
}

// AlibabaLstVendingOrderUpdateAPIResponseModel is 自动售货机订单物流信息回传 成功返回结果
type AlibabaLstVendingOrderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrorNo string `json:"error_no,omitempty" xml:"error_no,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 请求是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 是否执行了更新操作
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingOrderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorNo = ""
	m.ErrorMessage = ""
	m.Succ = false
	m.Module = false
}

var poolAlibabaLstVendingOrderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingOrderUpdateAPIResponse)
	},
}

// GetAlibabaLstVendingOrderUpdateAPIResponse 从 sync.Pool 获取 AlibabaLstVendingOrderUpdateAPIResponse
func GetAlibabaLstVendingOrderUpdateAPIResponse() *AlibabaLstVendingOrderUpdateAPIResponse {
	return poolAlibabaLstVendingOrderUpdateAPIResponse.Get().(*AlibabaLstVendingOrderUpdateAPIResponse)
}

// ReleaseAlibabaLstVendingOrderUpdateAPIResponse 将 AlibabaLstVendingOrderUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingOrderUpdateAPIResponse(v *AlibabaLstVendingOrderUpdateAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingOrderUpdateAPIResponse.Put(v)
}
