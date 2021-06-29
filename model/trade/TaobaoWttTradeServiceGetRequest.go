package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取网厅号卡垂直标信息 API请求
taobao.wtt.trade.service.get

查询网厅订单信息
*/
type TaobaoWttTradeServiceGetRequest struct {
    model.Params
    // 订单ID
    bizOrder   int64
}

// 初始化TaobaoWttTradeServiceGetRequest对象
func NewTaobaoWttTradeServiceGetRequest() *TaobaoWttTradeServiceGetRequest{
    return &TaobaoWttTradeServiceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWttTradeServiceGetRequest) GetApiMethodName() string {
    return "taobao.wtt.trade.service.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWttTradeServiceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrder Setter
// 订单ID
func (r *TaobaoWttTradeServiceGetRequest) SetBizOrder(bizOrder int64) error {
    r.bizOrder = bizOrder
    r.Set("biz_order", bizOrder)
    return nil
}

// BizOrder Getter
func (r TaobaoWttTradeServiceGetRequest) GetBizOrder() int64 {
    return r.bizOrder
}
