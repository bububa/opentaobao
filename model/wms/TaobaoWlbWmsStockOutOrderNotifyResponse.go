package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单通知 APIResponse
taobao.wlb.wms.stock.out.order.notify

出库单通知
*/
type TaobaoWlbWmsStockOutOrderNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_stock_out_order_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"