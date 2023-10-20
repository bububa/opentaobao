package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsCreatewarehouseorderAPIResponse 创建线上物流订单 API返回值
// aliexpress.logistics.createwarehouseorder
//
// 创建线上发货物流订单
type AliexpressLogisticsCreatewarehouseorderAPIResponse struct {
	model.CommonResponse
	AliexpressLogisticsCreatewarehouseorderAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLogisticsCreatewarehouseorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLogisticsCreatewarehouseorderAPIResponseModel).Reset()
}

// AliexpressLogisticsCreatewarehouseorderAPIResponseModel is 创建线上物流订单 成功返回结果
type AliexpressLogisticsCreatewarehouseorderAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_logistics_createwarehouseorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AeopWlCreateWarehouseOrderResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLogisticsCreatewarehouseorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
	m.ResultSuccess = false
}

var poolAliexpressLogisticsCreatewarehouseorderAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLogisticsCreatewarehouseorderAPIResponse)
	},
}

// GetAliexpressLogisticsCreatewarehouseorderAPIResponse 从 sync.Pool 获取 AliexpressLogisticsCreatewarehouseorderAPIResponse
func GetAliexpressLogisticsCreatewarehouseorderAPIResponse() *AliexpressLogisticsCreatewarehouseorderAPIResponse {
	return poolAliexpressLogisticsCreatewarehouseorderAPIResponse.Get().(*AliexpressLogisticsCreatewarehouseorderAPIResponse)
}

// ReleaseAliexpressLogisticsCreatewarehouseorderAPIResponse 将 AliexpressLogisticsCreatewarehouseorderAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLogisticsCreatewarehouseorderAPIResponse(v *AliexpressLogisticsCreatewarehouseorderAPIResponse) {
	v.Reset()
	poolAliexpressLogisticsCreatewarehouseorderAPIResponse.Put(v)
}
