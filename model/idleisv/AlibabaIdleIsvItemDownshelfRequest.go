package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品下架 APIRequest
alibaba.idle.isv.item.downshelf

供外部服务商ISV进行闲鱼商品下架操作
*/
type AlibabaIdleIsvItemDownshelfRequest struct {
    model.Params

    // 返回结果result
    param   *IdleItemBaseApiDo 

}

func NewAlibabaIdleIsvItemDownshelfRequest() *AlibabaIdleIsvItemDownshelfRequest{
    return &AlibabaIdleIsvItemDownshelfRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvItemDownshelfRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.downshelf"
}

func (r AlibabaIdleIsvItemDownshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvItemDownshelfRequest) SetParam(param *IdleItemBaseApiDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleIsvItemDownshelfRequest) GetParam() *IdleItemBaseApiDo {
    return r.param
}

