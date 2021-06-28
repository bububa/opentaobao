package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
存送业务查询奖品信息 APIResponse
alibaba.alicom.wtt.opentrade.getgiftdetails

话费宝充值送查询奖品信息
*/
type AlibabaAlicomWttOpentradeGetgiftdetailsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alicom_wtt_opentrade_getgiftdetails_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopResultDto `json:"result,omitempty" xml:"