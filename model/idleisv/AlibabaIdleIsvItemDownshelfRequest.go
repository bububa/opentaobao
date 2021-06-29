package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品下架 API请求
alibaba.idle.isv.item.downshelf

供外部服务商ISV进行闲鱼商品下架操作
*/
type AlibabaIdleIsvItemDownshelfRequest struct {
    model.Params
    // 返回结果result
    _param   *IdleItemBaseApiDO
}

// 初始化AlibabaIdleIsvItemDownshelfRequest对象
func NewAlibabaIdleIsvItemDownshelfRequest() *AlibabaIdleIsvItemDownshelfRequest{
    return &AlibabaIdleIsvItemDownshelfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemDownshelfRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.downshelf"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemDownshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 返回结果result
func (r *AlibabaIdleIsvItemDownshelfRequest) SetParam(_param *IdleItemBaseApiDO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaIdleIsvItemDownshelfRequest) GetParam() *IdleItemBaseApiDO {
    return r._param
}
