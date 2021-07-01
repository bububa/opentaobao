package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsStockOutOrderNotifyAPIResponse
出库单通知 API返回值
taobao.wlb.wms.stock.out.order.notify

出库单通知 */
type TaobaoWlbWmsStockOutOrderNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsStockOutOrderNotifyAPIResponseModel
}

// TaobaoWlbWmsStockOutOrderNotifyAPIResponseModel is 出库单通知 成功返回结果
type TaobaoWlbWmsStockOutOrderNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_stock_out_order_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// 是否成功
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
	// 错误详细
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// 仓储订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
}
