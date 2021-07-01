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
type AlibabaFundplatformAccountChargeNotifyAPIRequest struct {
    model.Params
    // 入参对象
    _request   *AlibabaFundplatformAccountChargeNotifyStruct
}

// 初始化AlibabaFundplatformAccountChargeNotifyAPIRequest对象
func NewAlibabaFundplatformAccountChargeNotifyRequest() *AlibabaFundplatformAccountChargeNotifyAPIRequest{
    return &AlibabaFundplatformAccountChargeNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.charge.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 入参对象
func (r *AlibabaFundplatformAccountChargeNotifyAPIRequest) SetRequest(_request *AlibabaFundplatformAccountChargeNotifyStruct) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r AlibabaFundplatformAccountChargeNotifyAPIRequest) GetRequest() *AlibabaFundplatformAccountChargeNotifyStruct {
    return r._request
}
