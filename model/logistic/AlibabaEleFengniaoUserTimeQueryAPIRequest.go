package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
蜂鸟询用户T API请求
alibaba.ele.fengniao.user.time.query

蜂鸟询用户T
*/
type AlibabaEleFengniaoUserTimeQueryAPIRequest struct {
    model.Params
    // 询T入参
    _param   *PredictDeliveryTimeParam
}

// 初始化AlibabaEleFengniaoUserTimeQueryAPIRequest对象
func NewAlibabaEleFengniaoUserTimeQueryRequest() *AlibabaEleFengniaoUserTimeQueryAPIRequest{
    return &AlibabaEleFengniaoUserTimeQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoUserTimeQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.user.time.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoUserTimeQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 询T入参
func (r *AlibabaEleFengniaoUserTimeQueryAPIRequest) SetParam(_param *PredictDeliveryTimeParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoUserTimeQueryAPIRequest) GetParam() *PredictDeliveryTimeParam {
    return r._param
}
