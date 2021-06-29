package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品查询 APIRequest
alibaba.idle.isv.item.query

服务商ISV闲鱼商品查询
*/
type AlibabaIdleIsvItemQueryRequest struct {
    model.Params

    // 商品查询参数
    param   *IdleItemBaseApiDo 

}

func NewAlibabaIdleIsvItemQueryRequest() *AlibabaIdleIsvItemQueryRequest{
    return &AlibabaIdleIsvItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvItemQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.query"
}

func (r AlibabaIdleIsvItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvItemQueryRequest) SetParam(param *IdleItemBaseApiDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleIsvItemQueryRequest) GetParam() *IdleItemBaseApiDo {
    return r.param
}

