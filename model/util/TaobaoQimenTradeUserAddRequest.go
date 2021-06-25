package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加奇门订单链路用户 APIRequest
taobao.qimen.trade.user.add

添加奇门订单链路用户
*/
type TaobaoQimenTradeUserAddRequest struct {
    model.Params

    // 商家备注
    memo   string 

}

func NewTaobaoQimenTradeUserAddRequest() *TaobaoQimenTradeUserAddRequest{
    return &TaobaoQimenTradeUserAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTradeUserAddRequest) GetApiMethodName() string {
    return "taobao.qimen.trade.user.add"
}

func (r TaobaoQimenTradeUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenTradeUserAddRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoQimenTradeUserAddRequest) GetMemo() string {
    return r.memo
}

