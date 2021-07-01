package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网厅垂直信息查询接口 API请求
taobao.trade.wtvertical.get

网厅订单垂直信息的查询
*/
type TaobaoTradeWtverticalGetAPIRequest struct {
    model.Params
    // 主订单列表,用“，”分隔tid的字符串,最大列表长度为15
    _tids   []int64
}

// 初始化TaobaoTradeWtverticalGetAPIRequest对象
func NewTaobaoTradeWtverticalGetRequest() *TaobaoTradeWtverticalGetAPIRequest{
    return &TaobaoTradeWtverticalGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeWtverticalGetAPIRequest) GetApiMethodName() string {
    return "taobao.trade.wtvertical.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeWtverticalGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tids Setter
// 主订单列表,用“，”分隔tid的字符串,最大列表长度为15
func (r *TaobaoTradeWtverticalGetAPIRequest) SetTids(_tids []int64) error {
    r._tids = _tids
    r.Set("tids", _tids)
    return nil
}

// Tids Getter
func (r TaobaoTradeWtverticalGetAPIRequest) GetTids() []int64 {
    return r._tids
}
