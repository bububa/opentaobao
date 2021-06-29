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
type AlibabaEleFengniaoUserTimeQueryRequest struct {
    model.Params
    // 询T入参
    _param   *PredictDeliveryTimeParam
}

// 初始化AlibabaEleFengniaoUserTimeQueryRequest对象
func NewAlibabaEleFengniaoUserTimeQueryRequest() *AlibabaEleFengniaoUserTimeQueryRequest{
    return &AlibabaEleFengniaoUserTimeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoUserTimeQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.user.time.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoUserTimeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 询T入参
func (r *AlibabaEleFengniaoUserTimeQueryRequest) SetParam(_param *PredictDeliveryTimeParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoUserTimeQueryRequest) GetParam() *PredictDeliveryTimeParam {
    return r._param
}
