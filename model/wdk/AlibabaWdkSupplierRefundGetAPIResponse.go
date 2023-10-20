package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierRefundGetAPIResponse 五道口按订单号批量查询供应商退款单 API返回值
// alibaba.wdk.supplier.refund.get
//
// 五道口按订单号批量查询供应商退款单
type AlibabaWdkSupplierRefundGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSupplierRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSupplierRefundGetAPIResponseModel).Reset()
}

// AlibabaWdkSupplierRefundGetAPIResponseModel is 五道口按订单号批量查询供应商退款单 成功返回结果
type AlibabaWdkSupplierRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSupplierRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSupplierRefundGetAPIResponse)
	},
}

// GetAlibabaWdkSupplierRefundGetAPIResponse 从 sync.Pool 获取 AlibabaWdkSupplierRefundGetAPIResponse
func GetAlibabaWdkSupplierRefundGetAPIResponse() *AlibabaWdkSupplierRefundGetAPIResponse {
	return poolAlibabaWdkSupplierRefundGetAPIResponse.Get().(*AlibabaWdkSupplierRefundGetAPIResponse)
}

// ReleaseAlibabaWdkSupplierRefundGetAPIResponse 将 AlibabaWdkSupplierRefundGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSupplierRefundGetAPIResponse(v *AlibabaWdkSupplierRefundGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkSupplierRefundGetAPIResponse.Put(v)
}
