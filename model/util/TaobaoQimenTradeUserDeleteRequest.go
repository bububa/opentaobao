package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除奇门订单链路用户 API请求
taobao.qimen.trade.user.delete

删除奇门订单链路用户
*/
type TaobaoQimenTradeUserDeleteRequest struct {
    model.Params
}

// 初始化TaobaoQimenTradeUserDeleteRequest对象
func NewTaobaoQimenTradeUserDeleteRequest() *TaobaoQimenTradeUserDeleteRequest{
    return &TaobaoQimenTradeUserDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUserDeleteRequest) GetApiMethodName() string {
    return "taobao.qimen.trade.user.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
