package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询连包签约状态 APIRequest
youku.ott.pay.order.queryauthstate

查询CP用户连包商品签约状态
*/
type YoukuOttPayOrderQueryauthstateRequest struct {
    model.Params

    // 原始签约订单号
    originalCpOrderNo   string 

}

func NewYoukuOttPayOrderQueryauthstateRequest() *YoukuOttPayOrderQueryauthstateRequest{
    return &YoukuOttPayOrderQueryauthstateRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderQueryauthstateRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.queryauthstate"
}

func (r YoukuOttPayOrderQueryauthstateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderQueryauthstateRequest) SetOriginalCpOrderNo(originalCpOrderNo string) error {
    r.originalCpOrderNo = originalCpOrderNo
    r.Set("original_cp_order_no", originalCpOrderNo)
    return nil
}

func (r YoukuOttPayOrderQueryauthstateRequest) GetOriginalCpOrderNo() string {
    return r.originalCpOrderNo
}

