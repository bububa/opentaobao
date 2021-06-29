package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单基本信息 API请求
alibaba.ele.fengniao.order.query

查询订单基本信息
*/
type AlibabaEleFengniaoOrderQueryRequest struct {
    model.Params
    // 参数
    param   *Param
}

// 初始化AlibabaEleFengniaoOrderQueryRequest对象
func NewAlibabaEleFengniaoOrderQueryRequest() *AlibabaEleFengniaoOrderQueryRequest{
    return &AlibabaEleFengniaoOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数
func (r *AlibabaEleFengniaoOrderQueryRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoOrderQueryRequest) GetParam() *Param {
    return r.param
}
