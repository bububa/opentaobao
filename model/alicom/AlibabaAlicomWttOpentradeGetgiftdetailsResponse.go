package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
存送业务查询奖品信息 APIResponse
alibaba.alicom.wtt.opentrade.getgiftdetails

话费宝充值送查询奖品信息
*/
type AlibabaAlicomWttOpentradeGetgiftdetailsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlicomWttOpentradeGetgiftdetailsResponse `json:"alibaba_alicom_wtt_opentrade_getgiftdetails_response,omitempty"`
}

type AlibabaAlicomWttOpentradeGetgiftdetailsResponse struct {

    // 结果
    Result   *TopResultDto `json:"result,omitempty"`

}