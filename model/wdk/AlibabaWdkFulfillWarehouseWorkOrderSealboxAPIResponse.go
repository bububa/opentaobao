package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse 仓封箱回告 API返回值
// alibaba.wdk.fulfill.warehouse.work.order.sealbox
//
// 仓封箱回告箱与包裹的关系
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel).Reset()
}

// AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel is 仓封箱回告 成功返回结果
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_warehouse_work_order_sealbox_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败返回原因
	RespMessage string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
	// 成功失败码
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespMessage = ""
	m.RespCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse)
	},
}

// GetAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse
func GetAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse() *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse {
	return poolAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse.Get().(*AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse)
}

// ReleaseAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse 将 AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse(v *AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse.Put(v)
}
