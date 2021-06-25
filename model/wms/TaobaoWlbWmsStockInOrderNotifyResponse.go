package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
入库通知单 APIResponse
taobao.wlb.wms.stock.in.order.notify

入库通知单
*/
type TaobaoWlbWmsStockInOrderNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsStockInOrderNotifyResponse `json:"taobao_wlb_wms_stock_in_order_notify_response,omitempty"`
}

type TaobaoWlbWmsStockInOrderNotifyResponse struct {

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误详细
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 仓储订单编码
    OrderCode   string `json:"order_code,omitempty"`

}
