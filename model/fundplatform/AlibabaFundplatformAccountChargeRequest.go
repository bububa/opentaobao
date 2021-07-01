package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资金平台余额账户充值 API请求
alibaba.fundplatform.account.charge

资金平台余额账户充值【创建账户&返回付款URL】
*/
type AlibabaFundplatformAccountChargeAPIRequest struct {
    model.Params
    // 用户ID
    _paramLong   int64
    // 入参对象
    _paramChargeRequest   *ChargeRequest
}

// 初始化AlibabaFundplatformAccountChargeAPIRequest对象
func NewAlibabaFundplatformAccountChargeRequest() *AlibabaFundplatformAccountChargeAPIRequest{
    return &AlibabaFundplatformAccountChargeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountChargeAPIRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.charge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountChargeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamLong Setter
// 用户ID
func (r *AlibabaFundplatformAccountChargeAPIRequest) SetParamLong(_paramLong int64) error {
    r._paramLong = _paramLong
    r.Set("param_long", _paramLong)
    return nil
}

// ParamLong Getter
func (r AlibabaFundplatformAccountChargeAPIRequest) GetParamLong() int64 {
    return r._paramLong
}
// ParamChargeRequest Setter
// 入参对象
func (r *AlibabaFundplatformAccountChargeAPIRequest) SetParamChargeRequest(_paramChargeRequest *ChargeRequest) error {
    r._paramChargeRequest = _paramChargeRequest
    r.Set("param_charge_request", _paramChargeRequest)
    return nil
}

// ParamChargeRequest Getter
func (r AlibabaFundplatformAccountChargeAPIRequest) GetParamChargeRequest() *ChargeRequest {
    return r._paramChargeRequest
}
