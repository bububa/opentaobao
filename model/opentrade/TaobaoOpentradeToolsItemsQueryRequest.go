package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易开放获取商品绑定信息 API请求
taobao.opentrade.tools.items.query

交易开放获取商品绑定信息
*/
type TaobaoOpentradeToolsItemsQueryRequest struct {
    model.Params
    // 交易开放C端小程序ID
    _miniappId   int64
}

// 初始化TaobaoOpentradeToolsItemsQueryRequest对象
func NewTaobaoOpentradeToolsItemsQueryRequest() *TaobaoOpentradeToolsItemsQueryRequest{
    return &TaobaoOpentradeToolsItemsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeToolsItemsQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.tools.items.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeToolsItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniappId Setter
// 交易开放C端小程序ID
func (r *TaobaoOpentradeToolsItemsQueryRequest) SetMiniappId(_miniappId int64) error {
    r._miniappId = _miniappId
    r.Set("miniapp_id", _miniappId)
    return nil
}

// MiniappId Getter
func (r TaobaoOpentradeToolsItemsQueryRequest) GetMiniappId() int64 {
    return r._miniappId
}
