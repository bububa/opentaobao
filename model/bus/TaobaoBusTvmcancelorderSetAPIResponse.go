package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmcancelorderSetAPIResponse 线下自助机未付款取消订单 API返回值
// taobao.bus.tvmcancelorder.set
//
// 自助机汽车票未付款取消订单
type TaobaoBusTvmcancelorderSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTvmcancelorderSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusTvmcancelorderSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusTvmcancelorderSetAPIResponseModel).Reset()
}

// TaobaoBusTvmcancelorderSetAPIResponseModel is 线下自助机未付款取消订单 成功返回结果
type TaobaoBusTvmcancelorderSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmcancelorder_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码  ORDER_NOT_FOUND
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述  订单无法查询
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// true代表成功 false 代表失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusTvmcancelorderSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusTvmcancelorderSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusTvmcancelorderSetAPIResponse)
	},
}

// GetTaobaoBusTvmcancelorderSetAPIResponse 从 sync.Pool 获取 TaobaoBusTvmcancelorderSetAPIResponse
func GetTaobaoBusTvmcancelorderSetAPIResponse() *TaobaoBusTvmcancelorderSetAPIResponse {
	return poolTaobaoBusTvmcancelorderSetAPIResponse.Get().(*TaobaoBusTvmcancelorderSetAPIResponse)
}

// ReleaseTaobaoBusTvmcancelorderSetAPIResponse 将 TaobaoBusTvmcancelorderSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusTvmcancelorderSetAPIResponse(v *TaobaoBusTvmcancelorderSetAPIResponse) {
	v.Reset()
	poolTaobaoBusTvmcancelorderSetAPIResponse.Put(v)
}
