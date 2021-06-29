package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单获取商品绑定信息 API请求
taobao.opentrade.special.items.query

专属下单获取商品绑定信息
*/
type TaobaoOpentradeSpecialItemsQueryRequest struct {
    model.Params
    // 绑定专属下单场景的C端小程序ID
    _miniappId   int64
}

// 初始化TaobaoOpentradeSpecialItemsQueryRequest对象
func NewTaobaoOpentradeSpecialItemsQueryRequest() *TaobaoOpentradeSpecialItemsQueryRequest{
    return &TaobaoOpentradeSpecialItemsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.items.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsQueryRequest) SetMiniappId(_miniappId int64) error {
    r._miniappId = _miniappId
    r.Set("miniapp_id", _miniappId)
    return nil
}

// MiniappId Getter
func (r TaobaoOpentradeSpecialItemsQueryRequest) GetMiniappId() int64 {
    return r._miniappId
}
