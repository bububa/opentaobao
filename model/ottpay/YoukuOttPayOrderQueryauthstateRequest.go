package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询连包签约状态 API请求
youku.ott.pay.order.queryauthstate

查询CP用户连包商品签约状态
*/
type YoukuOttPayOrderQueryauthstateRequest struct {
    model.Params
    // 原始签约订单号
    originalCpOrderNo   string
}

// 初始化YoukuOttPayOrderQueryauthstateRequest对象
func NewYoukuOttPayOrderQueryauthstateRequest() *YoukuOttPayOrderQueryauthstateRequest{
    return &YoukuOttPayOrderQueryauthstateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryauthstateRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.queryauthstate"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryauthstateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OriginalCpOrderNo Setter
// 原始签约订单号
func (r *YoukuOttPayOrderQueryauthstateRequest) SetOriginalCpOrderNo(originalCpOrderNo string) error {
    r.originalCpOrderNo = originalCpOrderNo
    r.Set("original_cp_order_no", originalCpOrderNo)
    return nil
}

// OriginalCpOrderNo Getter
func (r YoukuOttPayOrderQueryauthstateRequest) GetOriginalCpOrderNo() string {
    return r.originalCpOrderNo
}
