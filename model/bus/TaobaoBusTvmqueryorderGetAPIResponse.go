package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmqueryorderGetAPIResponse 线下自助机查询订单信息 API返回值
// taobao.bus.tvmqueryorder.get
//
// 查询订单详情
type TaobaoBusTvmqueryorderGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTvmqueryorderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusTvmqueryorderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusTvmqueryorderGetAPIResponseModel).Reset()
}

// TaobaoBusTvmqueryorderGetAPIResponseModel is 线下自助机查询订单信息 成功返回结果
type TaobaoBusTvmqueryorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmqueryorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// tvmBusOrderLineInfo
	TvmBusOrderLineInfo *TvmBusOrderLineInfo `json:"tvm_bus_order_line_info,omitempty" xml:"tvm_bus_order_line_info,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusTvmqueryorderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.TvmBusOrderLineInfo = nil
	m.IsSuccess = false
}

var poolTaobaoBusTvmqueryorderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusTvmqueryorderGetAPIResponse)
	},
}

// GetTaobaoBusTvmqueryorderGetAPIResponse 从 sync.Pool 获取 TaobaoBusTvmqueryorderGetAPIResponse
func GetTaobaoBusTvmqueryorderGetAPIResponse() *TaobaoBusTvmqueryorderGetAPIResponse {
	return poolTaobaoBusTvmqueryorderGetAPIResponse.Get().(*TaobaoBusTvmqueryorderGetAPIResponse)
}

// ReleaseTaobaoBusTvmqueryorderGetAPIResponse 将 TaobaoBusTvmqueryorderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusTvmqueryorderGetAPIResponse(v *TaobaoBusTvmqueryorderGetAPIResponse) {
	v.Reset()
	poolTaobaoBusTvmqueryorderGetAPIResponse.Put(v)
}
