package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送订单 API请求
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送
*/
type AlibabaEleFengniaoOrderPushAPIRequest struct {
    model.Params
    // 参数param
    _param   *Param
}

// 初始化AlibabaEleFengniaoOrderPushAPIRequest对象
func NewAlibabaEleFengniaoOrderPushRequest() *AlibabaEleFengniaoOrderPushAPIRequest{
    return &AlibabaEleFengniaoOrderPushAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.order.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数param
func (r *AlibabaEleFengniaoOrderPushAPIRequest) SetParam(_param *Param) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetParam() *Param {
    return r._param
}
