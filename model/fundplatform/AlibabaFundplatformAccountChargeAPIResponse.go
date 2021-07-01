package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资金平台余额账户充值 API返回值 
alibaba.fundplatform.account.charge

资金平台余额账户充值【创建账户&返回付款URL】
*/
type AlibabaFundplatformAccountChargeAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformAccountChargeAPIResponseModel
}

// 资金平台余额账户充值 成功返回结果
type AlibabaFundplatformAccountChargeAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_account_charge_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
