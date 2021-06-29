package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户充值成功通知 API请求
alibaba.fundplatform.account.charge.notify

通知外部业务方充值成功
*/
type AlibabaFundplatformAccountChargeNotifyRequest struct {
    model.Params
    // 入参对象
    _request   *AlibabaFundplatformAccountChargeNotifyStruct
}

// 初始化AlibabaFundplatformAccountChargeNotifyRequest对象
func NewAlibabaFundplatformAccountChargeNotifyRequest() *AlibabaFundplatformAccountChargeNotifyRequest{
    return &AlibabaFundplatformAccountChargeNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountChargeNotifyRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.charge.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountChargeNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 入参对象
func (r *AlibabaFundplatformAccountChargeNotifyRequest) SetRequest(_request *AlibabaFundplatformAccountChargeNotifyStruct) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r AlibabaFundplatformAccountChargeNotifyRequest) GetRequest() *AlibabaFundplatformAccountChargeNotifyStruct {
    return r._request
}
