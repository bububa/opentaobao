package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否为新人 APIRequest
alibaba.taobao.wt.user.crowd

增加isv接口
根据入参手机号判断是否为新人类型
*/
type AlibabaTaobaoWtUserCrowdRequest struct {
    model.Params

    // 手机号
    phone   int64 

}

func NewAlibabaTaobaoWtUserCrowdRequest() *AlibabaTaobaoWtUserCrowdRequest{
    return &AlibabaTaobaoWtUserCrowdRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTaobaoWtUserCrowdRequest) GetApiMethodName() string {
    return "alibaba.taobao.wt.user.crowd"
}

func (r AlibabaTaobaoWtUserCrowdRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTaobaoWtUserCrowdRequest) SetPhone(phone int64) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaTaobaoWtUserCrowdRequest) GetPhone() int64 {
    return r.phone
}

