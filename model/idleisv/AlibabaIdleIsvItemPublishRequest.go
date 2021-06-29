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
type AlibabaIdleIsvItemPublishRequest struct {
    model.Params
    // 商品数据参数
    itemParam   *IdleItemApiDo
}

// 初始化AlibabaIdleIsvItemPublishRequest对象
func NewAlibabaIdleIsvItemPublishRequest() *AlibabaIdleIsvItemPublishRequest{
    return &AlibabaIdleIsvItemPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemPublishRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.item.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemParam Setter
// 商品数据参数
func (r *AlibabaIdleIsvItemPublishRequest) SetItemParam(itemParam *IdleItemApiDo) error {
    r.itemParam = itemParam
    r.Set("item_param", itemParam)
    return nil
}

// ItemParam Getter
func (r AlibabaIdleIsvItemPublishRequest) GetItemParam() *IdleItemApiDo {
    return r.itemParam
}
