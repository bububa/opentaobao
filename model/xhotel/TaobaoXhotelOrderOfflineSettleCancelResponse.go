package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住取消结账专用接口 APIResponse
taobao.xhotel.order.offline.settle.cancel

线下信用住取消结账专用接口
*/
type TaobaoXhotelOrderOfflineSettleCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_order_offline_settle_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回信息
    
    Result   string `json:"result,omitempty" xml:"