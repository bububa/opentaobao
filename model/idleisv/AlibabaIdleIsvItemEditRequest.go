package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品编辑 API请求
alibaba.idle.isv.item.edit

服务商ISV闲鱼商品编辑操作
*/
type AlibabaIdleIsvItemEditRequest struct {
    model.Params
    // 商品数据参数
    _param   *IdleItemApiDo
}

// 初始化AlibabaIdleIsvItemEditRequest对象
func NewAlibabaIdleIsvItemEditRequest() *AlibabaIdleIsvItemEditRequest{
    return &AlibabaIdleIsvItemEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemEditRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.edit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 商品数据参数
func (r *AlibabaIdleIsvItemEditRequest) SetParam(_param *IdleItemApiDo) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaIdleIsvItemEditRequest) GetParam() *IdleItemApiDo {
    return r._param
}
