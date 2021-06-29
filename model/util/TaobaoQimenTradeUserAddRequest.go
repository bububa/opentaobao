package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加奇门订单链路用户 API请求
taobao.qimen.trade.user.add

添加奇门订单链路用户
*/
type TaobaoQimenTradeUserAddRequest struct {
    model.Params
    // 商家备注
    memo   string
}

// 初始化TaobaoQimenTradeUserAddRequest对象
func NewTaobaoQimenTradeUserAddRequest() *TaobaoQimenTradeUserAddRequest{
    return &TaobaoQimenTradeUserAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUserAddRequest) GetApiMethodName() string {
    return "taobao.qimen.trade.user.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Memo Setter
// 商家备注
func (r *TaobaoQimenTradeUserAddRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r TaobaoQimenTradeUserAddRequest) GetMemo() string {
    return r.memo
}
