package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierRefundListAPIResponse 五道口按供应商拉取退款单 API返回值
// alibaba.wdk.supplier.refund.list
//
// 五道口按供应商拉取退款单
type AlibabaWdkSupplierRefundListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSupplierRefundListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierRefundListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSupplierRefundListAPIResponseModel).Reset()
}

// AlibabaWdkSupplierRefundListAPIResponseModel is 五道口按供应商拉取退款单 成功返回结果
type AlibabaWdkSupplierRefundListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_refund_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierRefundListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSupplierRefundListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSupplierRefundListAPIResponse)
	},
}

// GetAlibabaWdkSupplierRefundListAPIResponse 从 sync.Pool 获取 AlibabaWdkSupplierRefundListAPIResponse
func GetAlibabaWdkSupplierRefundListAPIResponse() *AlibabaWdkSupplierRefundListAPIResponse {
	return poolAlibabaWdkSupplierRefundListAPIResponse.Get().(*AlibabaWdkSupplierRefundListAPIResponse)
}

// ReleaseAlibabaWdkSupplierRefundListAPIResponse 将 AlibabaWdkSupplierRefundListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSupplierRefundListAPIResponse(v *AlibabaWdkSupplierRefundListAPIResponse) {
	v.Reset()
	poolAlibabaWdkSupplierRefundListAPIResponse.Put(v)
}
