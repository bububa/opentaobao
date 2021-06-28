package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消物流订单接口 APIResponse
taobao.logistics.online.cancel

调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
*/
type TaobaoLogisticsOnlineCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"logistics_online_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功与失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"