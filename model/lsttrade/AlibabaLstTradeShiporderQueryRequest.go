package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商数据开放--发货单接口 APIRequest
alibaba.lst.trade.shiporder.query

供应商数据开放--发货单接口
*/
type AlibabaLstTradeShiporderQueryRequest struct {
    model.Params

    // 入参
    paramLstShipOrderQuery   *LstShipOrderQuery 

}

func NewAlibabaLstTradeShiporderQueryRequest() *AlibabaLstTradeShiporderQueryRequest{
    return &AlibabaLstTradeShiporderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeShiporderQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.shiporder.query"
}

func (r AlibabaLstTradeShiporderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeShiporderQueryRequest) SetParamLstShipOrderQuery(paramLstShipOrderQuery *LstShipOrderQuery) error {
    r.paramLstShipOrderQuery = paramLstShipOrderQuery
    r.Set("param_lst_ship_order_query", paramLstShipOrderQuery)
    return nil
}

func (r AlibabaLstTradeShiporderQueryRequest) GetParamLstShipOrderQuery() *LstShipOrderQuery {
    return r.paramLstShipOrderQuery
}

