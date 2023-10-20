package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse 退仓结果回传 API返回值
// alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed
//
// 退货入仓结果回传
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponseModel).Reset()
}

// AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponseModel is 退仓结果回传 成功返回结果
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_bill_return_warehouse_on_task_status_changed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// fulfillVoidResult
	Result *FulfillVoidResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse)
	},
}

// GetAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse
func GetAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse() *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse {
	return poolAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse.Get().(*AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse)
}

// ReleaseAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse 将 AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse(v *AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse.Put(v)
}
