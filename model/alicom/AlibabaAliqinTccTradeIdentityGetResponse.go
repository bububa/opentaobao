package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
运营商获得用户身份信息 APIResponse
alibaba.aliqin.tcc.trade.identity.get

天猫网厅运营商官方旗舰店获取用户身份信息
*/
type AlibabaAliqinTccTradeIdentityGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinTccTradeIdentityGetResponse `json:"alibaba_aliqin_tcc_trade_identity_get_response,omitempty"`
}

type AlibabaAliqinTccTradeIdentityGetResponse struct {

    // 返回身份信息
    Result   *IdentityInfo `json:"result,omitempty"`

}
