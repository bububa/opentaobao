package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库通知单 APIResponse
taobao.wlb.wms.stock.in.order.notify

入库通知单
*/
type TaobaoWlbWmsStockInOrderNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_stock_in_order_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"