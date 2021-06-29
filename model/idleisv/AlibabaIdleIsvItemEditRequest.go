package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品编辑 APIRequest
alibaba.idle.isv.item.edit

服务商ISV闲鱼商品编辑操作
*/
type AlibabaIdleIsvItemEditRequest struct {
    model.Params

    // 商品数据参数
    param   *IdleItemApiDo 

}

func NewAlibabaIdleIsvItemEditRequest() *AlibabaIdleIsvItemEditRequest{
    return &AlibabaIdleIsvItemEditRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleIsvItemEditRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.edit"
}

func (r AlibabaIdleIsvItemEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleIsvItemEditRequest) SetParam(param *IdleItemApiDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaIdleIsvItemEditRequest) GetParam() *IdleItemApiDo {
    return r.param
}

