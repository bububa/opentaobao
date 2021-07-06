package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsOrderCancelNotifyAPIResponse 单据取消接口 API返回值
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
type TaobaoWlbWmsOrderCancelNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsOrderCancelNotifyAPIResponseModel
}

// TaobaoWlbWmsOrderCancelNotifyAPIResponseModel is 单据取消接口 成功返回结果
type TaobaoWlbWmsOrderCancelNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_order_cancel_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 错误详细
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}
