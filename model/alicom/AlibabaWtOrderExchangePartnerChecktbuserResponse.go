package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
积分兑换校验淘宝账号是否存在 APIResponse
alibaba.wt.order.exchange.partner.checktbuser

积分兑换之前校验淘宝账号是否存在
*/
type AlibabaWtOrderExchangePartnerChecktbuserAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wt_order_exchange_partner_checktbuser_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值，通过model的值true或者false来判断
    
    Model   bool `json:"model,omitempty" xml:"