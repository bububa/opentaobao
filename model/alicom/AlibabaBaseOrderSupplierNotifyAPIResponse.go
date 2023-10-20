package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaseOrderSupplierNotifyAPIResponse 阿里通信运营商信息回传 API返回值
// alibaba.base.order.supplier.notify
//
// 接收阿里通信流量运营商信息回传
type AlibabaBaseOrderSupplierNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaBaseOrderSupplierNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaseOrderSupplierNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaseOrderSupplierNotifyAPIResponseModel).Reset()
}

// AlibabaBaseOrderSupplierNotifyAPIResponseModel is 阿里通信运营商信息回传 成功返回结果
type AlibabaBaseOrderSupplierNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_base_order_supplier_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaseOrderSupplierNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaseOrderSupplierNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaseOrderSupplierNotifyAPIResponse)
	},
}

// GetAlibabaBaseOrderSupplierNotifyAPIResponse 从 sync.Pool 获取 AlibabaBaseOrderSupplierNotifyAPIResponse
func GetAlibabaBaseOrderSupplierNotifyAPIResponse() *AlibabaBaseOrderSupplierNotifyAPIResponse {
	return poolAlibabaBaseOrderSupplierNotifyAPIResponse.Get().(*AlibabaBaseOrderSupplierNotifyAPIResponse)
}

// ReleaseAlibabaBaseOrderSupplierNotifyAPIResponse 将 AlibabaBaseOrderSupplierNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaseOrderSupplierNotifyAPIResponse(v *AlibabaBaseOrderSupplierNotifyAPIResponse) {
	v.Reset()
	poolAlibabaBaseOrderSupplierNotifyAPIResponse.Put(v)
}
