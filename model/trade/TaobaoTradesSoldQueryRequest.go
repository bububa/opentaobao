package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号 API请求
taobao.trades.sold.query

根据收件人信息查询交易单号。
*/
type TaobaoTradesSoldQueryRequest struct {
    model.Params
    // 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
    queryList   []OrderQuery
}

// 初始化TaobaoTradesSoldQueryRequest对象
func NewTaobaoTradesSoldQueryRequest() *TaobaoTradesSoldQueryRequest{
    return &TaobaoTradesSoldQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradesSoldQueryRequest) GetApiMethodName() string {
    return "taobao.trades.sold.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradesSoldQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryList Setter
// 查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。
func (r *TaobaoTradesSoldQueryRequest) SetQueryList(queryList []OrderQuery) error {
    r.queryList = queryList
    r.Set("query_list", queryList)
    return nil
}

// QueryList Getter
func (r TaobaoTradesSoldQueryRequest) GetQueryList() []OrderQuery {
    return r.queryList
}
