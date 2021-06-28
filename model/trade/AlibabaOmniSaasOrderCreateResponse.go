package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单创建接口 APIResponse
alibaba.omni.saas.order.create

服务商利用现有的saas系统和阿里完成交易系统的对接
*/
type AlibabaOmniSaasOrderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_omni_saas_order_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // tradeNo
    
    TradeNo   int64 `json:"trade_no,omitempty" xml:"