package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资金平台余额账户充值 APIRequest
alibaba.fundplatform.account.charge

资金平台余额账户充值【创建账户&返回付款URL】
*/
type AlibabaFundplatformAccountChargeRequest struct {
    model.Params

    // 用户ID
    paramLong   int64 

    // 入参对象
    paramChargeRequest   *ChargeRequest 

}

func NewAlibabaFundplatformAccountChargeRequest() *AlibabaFundplatformAccountChargeRequest{
    return &AlibabaFundplatformAccountChargeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformAccountChargeRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.charge"
}

func (r AlibabaFundplatformAccountChargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformAccountChargeRequest) SetParamLong(paramLong int64) error {
    r.paramLong = paramLong
    r.Set("param_long", paramLong)
    return nil
}

func (r AlibabaFundplatformAccountChargeRequest) GetParamLong() int64 {
    return r.paramLong
}

func (r *AlibabaFundplatformAccountChargeRequest) SetParamChargeRequest(paramChargeRequest *ChargeRequest) error {
    r.paramChargeRequest = paramChargeRequest
    r.Set("param_charge_request", paramChargeRequest)
    return nil
}

func (r AlibabaFundplatformAccountChargeRequest) GetParamChargeRequest() *ChargeRequest {
    return r.paramChargeRequest
}

