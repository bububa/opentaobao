package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商停止发货 APIRequest
tmall.channel.trade.order.stop

供应商停止发货
*/
type TmallChannelTradeOrderStopRequest struct {
    model.Params

    // 主采购单号
    mainPurchaseOrderNo   int64 

    // 幂等单号
    requestNo   string 

}

func NewTmallChannelTradeOrderStopRequest() *TmallChannelTradeOrderStopRequest{
    return &TmallChannelTradeOrderStopRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeOrderStopRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.stop"
}

func (r TmallChannelTradeOrderStopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeOrderStopRequest) SetMainPurchaseOrderNo(mainPurchaseOrderNo int64) error {
    r.mainPurchaseOrderNo = mainPurchaseOrderNo
    r.Set("main_purchase_order_no", mainPurchaseOrderNo)
    return nil
}

func (r TmallChannelTradeOrderStopRequest) GetMainPurchaseOrderNo() int64 {
    return r.mainPurchaseOrderNo
}

func (r *TmallChannelTradeOrderStopRequest) SetRequestNo(requestNo string) error {
    r.requestNo = requestNo
    r.Set("request_no", requestNo)
    return nil
}

func (r TmallChannelTradeOrderStopRequest) GetRequestNo() string {
    return r.requestNo
}

