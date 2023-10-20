package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderCancelAPIResponse 取消物流宝订单 API返回值
// taobao.wlb.order.cancel
//
// 取消物流宝订单
type TaobaoWlbOrderCancelAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderCancelAPIResponseModel).Reset()
}

// TaobaoWlbOrderCancelAPIResponseModel is 取消物流宝订单 成功返回结果
type TaobaoWlbOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改时间，只有在取消成功的情况下，才可以做
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 错误编码列表
	ErrorCodeList string `json:"error_code_list,omitempty" xml:"error_code_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModifyTime = ""
	m.ErrorCodeList = ""
}

var poolTaobaoWlbOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderCancelAPIResponse)
	},
}

// GetTaobaoWlbOrderCancelAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderCancelAPIResponse
func GetTaobaoWlbOrderCancelAPIResponse() *TaobaoWlbOrderCancelAPIResponse {
	return poolTaobaoWlbOrderCancelAPIResponse.Get().(*TaobaoWlbOrderCancelAPIResponse)
}

// ReleaseTaobaoWlbOrderCancelAPIResponse 将 TaobaoWlbOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderCancelAPIResponse(v *TaobaoWlbOrderCancelAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderCancelAPIResponse.Put(v)
}
