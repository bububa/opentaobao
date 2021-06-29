package happytrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商渠道余额 APIResponse
alibaba.happytrip.taxi.provider.account.balance

查询不同供应商不同渠道账户余额
*/
type AlibabaHappytripTaxiProviderAccountBalanceAPIResponse struct {
    model.CommonResponse
    AlibabaHappytripTaxiProviderAccountBalanceResponse
}

type AlibabaHappytripTaxiProviderAccountBalanceResponse struct {
    XMLName xml.Name `xml:"alibaba_happytrip_taxi_provider_account_balance_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Errno   int64 `json:"errno,omitempty" xml:"errno,omitempty"`

    
    // 数据
    
    Data   *AlibabaHappytripTaxiProviderAccountBalanceData `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
}
