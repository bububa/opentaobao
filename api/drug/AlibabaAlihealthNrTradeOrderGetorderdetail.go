package drug

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drug"
)

/* 
根据订单id获取单条订单详情 
alibaba.alihealth.nr.trade.order.getorderdetail

阿里健康O2O，获取订单详情，修复组合商品价格精度问题
*/
func AlibabaAlihealthNrTradeOrderGetorderdetail(clt *core.SDKClient, req *drug.AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest, session string) (*drug.AlibabaAlihealthNrTradeOrderGetorderdetailAPIResponse, error) {
    var resp drug.AlibabaAlihealthNrTradeOrderGetorderdetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
