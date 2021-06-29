package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取健康id API请求
alibaba.alihealth.uic.userinfo.healthid.get

根据支付宝用户ID获取用户健康ID
*/
type AlibabaAlihealthUicUserinfoHealthidGetRequest struct {
    model.Params
    // 支付宝用户ID
    _alipayUserId   string
}

// 初始化AlibabaAlihealthUicUserinfoHealthidGetRequest对象
func NewAlibabaAlihealthUicUserinfoHealthidGetRequest() *AlibabaAlihealthUicUserinfoHealthidGetRequest{
    return &AlibabaAlihealthUicUserinfoHealthidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthUicUserinfoHealthidGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.uic.userinfo.healthid.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthUicUserinfoHealthidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝用户ID
func (r *AlibabaAlihealthUicUserinfoHealthidGetRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaAlihealthUicUserinfoHealthidGetRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
