package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmcreateorderSetAPIResponse 线下自助机创建订单 API返回值
// taobao.bus.tvmcreateorder.set
//
// 提供给汽车票线下自助机的创建订单使用
type TaobaoBusTvmcreateorderSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTvmcreateorderSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusTvmcreateorderSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusTvmcreateorderSetAPIResponseModel).Reset()
}

// TaobaoBusTvmcreateorderSetAPIResponseModel is 线下自助机创建订单 成功返回结果
type TaobaoBusTvmcreateorderSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmcreateorder_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alitripOrderId
	AlitripOrderId string `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`
	// errorCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusTvmcreateorderSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlitripOrderId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusTvmcreateorderSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusTvmcreateorderSetAPIResponse)
	},
}

// GetTaobaoBusTvmcreateorderSetAPIResponse 从 sync.Pool 获取 TaobaoBusTvmcreateorderSetAPIResponse
func GetTaobaoBusTvmcreateorderSetAPIResponse() *TaobaoBusTvmcreateorderSetAPIResponse {
	return poolTaobaoBusTvmcreateorderSetAPIResponse.Get().(*TaobaoBusTvmcreateorderSetAPIResponse)
}

// ReleaseTaobaoBusTvmcreateorderSetAPIResponse 将 TaobaoBusTvmcreateorderSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusTvmcreateorderSetAPIResponse(v *TaobaoBusTvmcreateorderSetAPIResponse) {
	v.Reset()
	poolTaobaoBusTvmcreateorderSetAPIResponse.Put(v)
}
