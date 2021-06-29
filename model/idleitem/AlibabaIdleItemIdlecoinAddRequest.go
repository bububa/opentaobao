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
type AlibabaIdleItemIdlecoinAddRequest struct {
    model.Params
    // 免费送商品数据
    idleCoinItemParam   *IdleCoinItemApiDto
}

// 初始化AlibabaIdleItemIdlecoinAddRequest对象
func NewAlibabaIdleItemIdlecoinAddRequest() *AlibabaIdleItemIdlecoinAddRequest{
    return &AlibabaIdleItemIdlecoinAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemIdlecoinAddRequest) GetApiMethodName() string {
    return "alibaba.idle.item.idlecoin.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemIdlecoinAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdleCoinItemParam Setter
// 免费送商品数据
func (r *AlibabaIdleItemIdlecoinAddRequest) SetIdleCoinItemParam(idleCoinItemParam *IdleCoinItemApiDto) error {
    r.idleCoinItemParam = idleCoinItemParam
    r.Set("idle_coin_item_param", idleCoinItemParam)
    return nil
}

// IdleCoinItemParam Getter
func (r AlibabaIdleItemIdlecoinAddRequest) GetIdleCoinItemParam() *IdleCoinItemApiDto {
    return r.idleCoinItemParam
}
