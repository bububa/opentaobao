package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通交易订单查询--品牌商视角 APIRequest
alibaba.lst.trade.order.get

根据订单id查询零售通交易订单
*/
type AlibabaLstTradeOrderGetRequest struct {
    model.Params

    // 主订单id
    mainOrderId   int64 

    // 子订单id
    subOrderId   int64 

}

func NewAlibabaLstTradeOrderGetRequest() *AlibabaLstTradeOrderGetRequest{
    return &AlibabaLstTradeOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeOrderGetRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.get"
}

func (r AlibabaLstTradeOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeOrderGetRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r AlibabaLstTradeOrderGetRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *AlibabaLstTradeOrderGetRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r AlibabaLstTradeOrderGetRequest) GetSubOrderId() int64 {
    return r.subOrderId
}

