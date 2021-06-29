package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资金平台余额账户充值 APIResponse
alibaba.fundplatform.account.charge

资金平台余额账户充值【创建账户&返回付款URL】
*/
type AlibabaFundplatformAccountChargeAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformAccountChargeResponse
}

type AlibabaFundplatformAccountChargeResponse struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_account_charge_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
