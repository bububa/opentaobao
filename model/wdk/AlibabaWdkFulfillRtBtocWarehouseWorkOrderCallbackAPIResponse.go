package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse 大润发B2C仓作业单回传接口 API返回值
// alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback
//
// 大润发B2C仓作业单回传接口
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponseModel).Reset()
}

// AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponseModel is 大润发B2C仓作业单回传接口 成功返回结果
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_rt_btoc_warehouse_work_order_callback_response"`
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
func (m *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespMessage = ""
	m.RespCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse)
	},
}

// GetAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse
func GetAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse() *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse {
	return poolAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse.Get().(*AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse)
}

// ReleaseAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse 将 AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse(v *AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse.Put(v)
}
