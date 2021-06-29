package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否为新人 API请求
alibaba.taobao.wt.user.crowd

增加isv接口
根据入参手机号判断是否为新人类型
*/
type AlibabaTaobaoWtUserCrowdRequest struct {
    model.Params
    // 手机号
    _phone   int64
}

// 初始化AlibabaTaobaoWtUserCrowdRequest对象
func NewAlibabaTaobaoWtUserCrowdRequest() *AlibabaTaobaoWtUserCrowdRequest{
    return &AlibabaTaobaoWtUserCrowdRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoWtUserCrowdRequest) GetApiMethodName() string {
    return "alibaba.taobao.wt.user.crowd"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoWtUserCrowdRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Phone Setter
// 手机号
func (r *AlibabaTaobaoWtUserCrowdRequest) SetPhone(_phone int64) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaTaobaoWtUserCrowdRequest) GetPhone() int64 {
    return r._phone
}
