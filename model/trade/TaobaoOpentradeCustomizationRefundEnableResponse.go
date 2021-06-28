package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定制订单设置允许仅退款 APIResponse
taobao.opentrade.customization.refund.enable

定制订单设置允许仅退款
*/
type TaobaoOpentradeCustomizationRefundEnableAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"opentrade_customization_refund_enable_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否设置成功
    
    Result   bool `json:"result,omitempty" xml:"