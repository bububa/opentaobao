package icbuassurance

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu信保账户信息 API返回值 
alibaba.icbu.trade.assurance.account.get

icbu交易信用保障开通状态&额度信息查询
*/
type AlibabaIcbuTradeAssuranceAccountGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuTradeAssuranceAccountGetResponse
}

// icbu信保账户信息 成功返回结果
type AlibabaIcbuTradeAssuranceAccountGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_trade_assurance_account_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // AssuranceAccountResult
    AssuranceAccountResult   *AssuranceAccountResult `json:"assurance_account_result,omitempty" xml:"assurance_account_result,omitempty"`
}
