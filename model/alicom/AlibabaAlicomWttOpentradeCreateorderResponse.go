package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
充值送活动下单接口 APIResponse
alibaba.alicom.wtt.opentrade.createorder

提供给话费宝创建淘宝订单
*/
type AlibabaAlicomWttOpentradeCreateorderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alicom_wtt_opentrade_createorder_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopResultDto `json:"result,omitempty" xml:"