package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品发布 API请求
alibaba.idle.isv.item.publish

服务商ISV闲鱼商品发布
*/
type AlibabaIdleIsvItemPublishAPIRequest struct {
    model.Params
    // 商品数据参数
    _itemParam   *IdleItemApiDO
}

// 初始化AlibabaIdleIsvItemPublishAPIRequest对象
func NewAlibabaIdleIsvItemPublishRequest() *AlibabaIdleIsvItemPublishAPIRequest{
    return &AlibabaIdleIsvItemPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemPublishAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemParam Setter
// 商品数据参数
func (r *AlibabaIdleIsvItemPublishAPIRequest) SetItemParam(_itemParam *IdleItemApiDO) error {
    r._itemParam = _itemParam
    r.Set("item_param", _itemParam)
    return nil
}

// ItemParam Getter
func (r AlibabaIdleIsvItemPublishAPIRequest) GetItemParam() *IdleItemApiDO {
    return r._itemParam
}
