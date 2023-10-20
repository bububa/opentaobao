package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierOrderListAPIResponse 五道口供应商维度正向订单拉取 API返回值
// alibaba.wdk.supplier.order.list
//
// 五道口供应商维度正向订单拉取
type AlibabaWdkSupplierOrderListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSupplierOrderListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierOrderListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSupplierOrderListAPIResponseModel).Reset()
}

// AlibabaWdkSupplierOrderListAPIResponseModel is 五道口供应商维度正向订单拉取 成功返回结果
type AlibabaWdkSupplierOrderListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderListSyncPagedResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierOrderListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSupplierOrderListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSupplierOrderListAPIResponse)
	},
}

// GetAlibabaWdkSupplierOrderListAPIResponse 从 sync.Pool 获取 AlibabaWdkSupplierOrderListAPIResponse
func GetAlibabaWdkSupplierOrderListAPIResponse() *AlibabaWdkSupplierOrderListAPIResponse {
	return poolAlibabaWdkSupplierOrderListAPIResponse.Get().(*AlibabaWdkSupplierOrderListAPIResponse)
}

// ReleaseAlibabaWdkSupplierOrderListAPIResponse 将 AlibabaWdkSupplierOrderListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSupplierOrderListAPIResponse(v *AlibabaWdkSupplierOrderListAPIResponse) {
	v.Reset()
	poolAlibabaWdkSupplierOrderListAPIResponse.Put(v)
}
