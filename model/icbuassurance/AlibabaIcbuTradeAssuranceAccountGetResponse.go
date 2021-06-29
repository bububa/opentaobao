package icbuassurance

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu信保账户信息 APIResponse
alibaba.icbu.trade.assurance.account.get

icbu交易信用保障开通状态&额度信息查询
*/
type AlibabaIcbuTradeAssuranceAccountGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuTradeAssuranceAccountGetResponse
}

type AlibabaIcbuTradeAssuranceAccountGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_trade_assurance_account_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // AssuranceAccountResult
    
    AssuranceAccountResult   *AssuranceAccountResult `json:"assurance_account_result,omitempty" xml:"assurance_account_result,omitempty"`

    
}
