package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单基本信息 APIRequest
alibaba.ele.fengniao.order.query

查询订单基本信息
*/
type AlibabaEleFengniaoOrderQueryRequest struct {
    model.Params

    // 参数
    param   *Param 

}

func NewAlibabaEleFengniaoOrderQueryRequest() *AlibabaEleFengniaoOrderQueryRequest{
    return &AlibabaEleFengniaoOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.order.query"
}

func (r AlibabaEleFengniaoOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoOrderQueryRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoOrderQueryRequest) GetParam() *Param {
    return r.param
}

