package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商数据开放--发货单接口 API请求
alibaba.lst.trade.shiporder.query

供应商数据开放--发货单接口
*/
type AlibabaLstTradeShiporderQueryRequest struct {
    model.Params
    // 入参
    _paramLstShipOrderQuery   *LstShipOrderQuery
}

// 初始化AlibabaLstTradeShiporderQueryRequest对象
func NewAlibabaLstTradeShiporderQueryRequest() *AlibabaLstTradeShiporderQueryRequest{
    return &AlibabaLstTradeShiporderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeShiporderQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.shiporder.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeShiporderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamLstShipOrderQuery Setter
// 入参
func (r *AlibabaLstTradeShiporderQueryRequest) SetParamLstShipOrderQuery(_paramLstShipOrderQuery *LstShipOrderQuery) error {
    r._paramLstShipOrderQuery = _paramLstShipOrderQuery
    r.Set("param_lst_ship_order_query", _paramLstShipOrderQuery)
    return nil
}

// ParamLstShipOrderQuery Getter
func (r AlibabaLstTradeShiporderQueryRequest) GetParamLstShipOrderQuery() *LstShipOrderQuery {
    return r._paramLstShipOrderQuery
}
