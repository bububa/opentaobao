package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse 标准化仓作业单回传接口 API返回值
// alibaba.wdk.fulfill.warehouse.work.order.callback
//
// 标准化仓作业单回传接口
type AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponseModel).Reset()
}

// AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponseModel is 标准化仓作业单回传接口 成功返回结果
type AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_warehouse_work_order_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应提示信息
	RespMessage string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
	// 响应code
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespMessage = ""
	m.RespCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse)
	},
}

// GetAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse
func GetAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse() *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse {
	return poolAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse.Get().(*AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse)
}

// ReleaseAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse 将 AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse(v *AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse.Put(v)
}
