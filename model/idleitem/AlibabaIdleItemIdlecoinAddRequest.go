package idleitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
免费送商品发送 API请求
alibaba.idle.item.idlecoin.add

免费送商品发布
*/
type AlibabaIdleItemIdlecoinAddAPIRequest struct {
    model.Params
    // 免费送商品数据
    _idleCoinItemParam   *IdleCoinItemApiDTO
}

// 初始化AlibabaIdleItemIdlecoinAddAPIRequest对象
func NewAlibabaIdleItemIdlecoinAddRequest() *AlibabaIdleItemIdlecoinAddAPIRequest{
    return &AlibabaIdleItemIdlecoinAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemIdlecoinAddAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.item.idlecoin.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemIdlecoinAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdleCoinItemParam Setter
// 免费送商品数据
func (r *AlibabaIdleItemIdlecoinAddAPIRequest) SetIdleCoinItemParam(_idleCoinItemParam *IdleCoinItemApiDTO) error {
    r._idleCoinItemParam = _idleCoinItemParam
    r.Set("idle_coin_item_param", _idleCoinItemParam)
    return nil
}

// IdleCoinItemParam Getter
func (r AlibabaIdleItemIdlecoinAddAPIRequest) GetIdleCoinItemParam() *IdleCoinItemApiDTO {
    return r._idleCoinItemParam
}
