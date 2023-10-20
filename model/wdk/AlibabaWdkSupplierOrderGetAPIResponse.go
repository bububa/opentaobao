package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierOrderGetAPIResponse 五道口按订单号批量查询供应商正向订单 API返回值
// alibaba.wdk.supplier.order.get
//
// 五道口按订单号批量查询供应商正向订单
type AlibabaWdkSupplierOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSupplierOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSupplierOrderGetAPIResponseModel).Reset()
}

// AlibabaWdkSupplierOrderGetAPIResponseModel is 五道口按订单号批量查询供应商正向订单 成功返回结果
type AlibabaWdkSupplierOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderListSyncPagedResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSupplierOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSupplierOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSupplierOrderGetAPIResponse)
	},
}

// GetAlibabaWdkSupplierOrderGetAPIResponse 从 sync.Pool 获取 AlibabaWdkSupplierOrderGetAPIResponse
func GetAlibabaWdkSupplierOrderGetAPIResponse() *AlibabaWdkSupplierOrderGetAPIResponse {
	return poolAlibabaWdkSupplierOrderGetAPIResponse.Get().(*AlibabaWdkSupplierOrderGetAPIResponse)
}

// ReleaseAlibabaWdkSupplierOrderGetAPIResponse 将 AlibabaWdkSupplierOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSupplierOrderGetAPIResponse(v *AlibabaWdkSupplierOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkSupplierOrderGetAPIResponse.Put(v)
}
