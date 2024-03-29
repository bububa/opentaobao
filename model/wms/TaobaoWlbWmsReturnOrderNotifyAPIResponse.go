package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsReturnOrderNotifyAPIResponse 销售退货通知 API返回值
// taobao.wlb.wms.return.order.notify
//
// 销售退货通知
type TaobaoWlbWmsReturnOrderNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsReturnOrderNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsReturnOrderNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsReturnOrderNotifyAPIResponseModel).Reset()
}

// TaobaoWlbWmsReturnOrderNotifyAPIResponseModel is 销售退货通知 成功返回结果
type TaobaoWlbWmsReturnOrderNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_return_order_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// LBX929829111
	StoreOrderCode string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 错误信息
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsReturnOrderNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CreateTime = ""
	m.StoreOrderCode = ""
	m.WlErrorCode = ""
	m.WlErrorMsg = ""
	m.WlSuccess = false
}

var poolTaobaoWlbWmsReturnOrderNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsReturnOrderNotifyAPIResponse)
	},
}

// GetTaobaoWlbWmsReturnOrderNotifyAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsReturnOrderNotifyAPIResponse
func GetTaobaoWlbWmsReturnOrderNotifyAPIResponse() *TaobaoWlbWmsReturnOrderNotifyAPIResponse {
	return poolTaobaoWlbWmsReturnOrderNotifyAPIResponse.Get().(*TaobaoWlbWmsReturnOrderNotifyAPIResponse)
}

// ReleaseTaobaoWlbWmsReturnOrderNotifyAPIResponse 将 TaobaoWlbWmsReturnOrderNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsReturnOrderNotifyAPIResponse(v *TaobaoWlbWmsReturnOrderNotifyAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsReturnOrderNotifyAPIResponse.Put(v)
}
