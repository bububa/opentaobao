package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网厅垂直信息查询接口 APIRequest
taobao.trade.wtvertical.get

网厅订单垂直信息的查询
*/
type TaobaoTradeWtverticalGetRequest struct {
    model.Params

    // 主订单列表,用“，”分隔tid的字符串,最大列表长度为15
    tids   []int64 

}

func NewTaobaoTradeWtverticalGetRequest() *TaobaoTradeWtverticalGetRequest{
    return &TaobaoTradeWtverticalGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeWtverticalGetRequest) GetApiMethodName() string {
    return "taobao.trade.wtvertical.get"
}

func (r TaobaoTradeWtverticalGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeWtverticalGetRequest) SetTids(tids []int64) error {
    r.tids = tids
    r.Set("tids", tids)
    return nil
}

func (r TaobaoTradeWtverticalGetRequest) GetTids() []int64 {
    return r.tids
}

