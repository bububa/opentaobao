package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderShipAPIResponse 闲鱼订单服务商物流发货 API返回值
// alibaba.idle.isv.order.ship
//
// 闲鱼开放平台服务商订单发货接口
type AlibabaIdleIsvOrderShipAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvOrderShipAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderShipAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvOrderShipAPIResponseModel).Reset()
}

// AlibabaIdleIsvOrderShipAPIResponseModel is 闲鱼订单服务商物流发货 成功返回结果
type AlibabaIdleIsvOrderShipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_order_ship_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultErrCode string `json:"result_err_code,omitempty" xml:"result_err_code,omitempty"`
	// 错误信息
	ResultErrMsg string `json:"result_err_msg,omitempty" xml:"result_err_msg,omitempty"`
	// 请求结果，是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
	// 业务逻辑结果，暂时不用
	ResultModule bool `json:"result_module,omitempty" xml:"result_module,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderShipAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultErrCode = ""
	m.ResultErrMsg = ""
	m.ResultSuccess = false
	m.ResultModule = false
}

var poolAlibabaIdleIsvOrderShipAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvOrderShipAPIResponse)
	},
}

// GetAlibabaIdleIsvOrderShipAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvOrderShipAPIResponse
func GetAlibabaIdleIsvOrderShipAPIResponse() *AlibabaIdleIsvOrderShipAPIResponse {
	return poolAlibabaIdleIsvOrderShipAPIResponse.Get().(*AlibabaIdleIsvOrderShipAPIResponse)
}

// ReleaseAlibabaIdleIsvOrderShipAPIResponse 将 AlibabaIdleIsvOrderShipAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvOrderShipAPIResponse(v *AlibabaIdleIsvOrderShipAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvOrderShipAPIResponse.Put(v)
}
