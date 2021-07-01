package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsConsignOrderNotifyAPIResponse
发货订单通知 API返回值
taobao.wlb.wms.consign.order.notify

发货订单通知 */
type TaobaoWlbWmsConsignOrderNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsConsignOrderNotifyAPIResponseModel
}

// TaobaoWlbWmsConsignOrderNotifyAPIResponseModel is 发货订单通知 成功返回结果
type TaobaoWlbWmsConsignOrderNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_consign_order_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 错误详细
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 系统自动生成
	ConsignOrderList []Consignorderlist `json:"consign_order_list,omitempty" xml:"consign_order_list>consignorderlist,omitempty"`
}
