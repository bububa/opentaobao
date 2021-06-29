package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品查询 API请求
alibaba.idle.isv.item.query

服务商ISV闲鱼商品查询
*/
type AlibabaIdleIsvItemQueryRequest struct {
    model.Params
    // 商品查询参数
    _param   *IdleItemBaseApiDO
}

// 初始化AlibabaIdleIsvItemQueryRequest对象
func NewAlibabaIdleIsvItemQueryRequest() *AlibabaIdleIsvItemQueryRequest{
    return &AlibabaIdleIsvItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 商品查询参数
func (r *AlibabaIdleIsvItemQueryRequest) SetParam(_param *IdleItemBaseApiDO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaIdleIsvItemQueryRequest) GetParam() *IdleItemBaseApiDO {
    return r._param
}
