package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单并发货 APIResponse
taobao.logistics.consign.order.createandsend

创建物流订单，并发货。
*/
type TaobaoLogisticsConsignOrderCreateandsendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"logistics_consign_order_createandsend_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果描述
    
    ResultDesc   string `json:"result_desc,omitempty" xml:"