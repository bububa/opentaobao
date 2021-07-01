package lsttrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lsttrade"
)

/* 
供应商数据开放--发货单接口 
alibaba.lst.trade.shiporder.query

供应商数据开放--发货单接口
*/
func AlibabaLstTradeShiporderQuery(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeShiporderQueryAPIRequest, session string) (*lsttrade.AlibabaLstTradeShiporderQueryAPIResponse, error) {
    var resp lsttrade.AlibabaLstTradeShiporderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
