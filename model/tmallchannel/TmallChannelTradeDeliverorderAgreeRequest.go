package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核通过发货确认 APIRequest
tmall.channel.trade.deliverorder.agree

供应商通过收货确认单
*/
type TmallChannelTradeDeliverorderAgreeRequest struct {
    model.Params

    // 发货单号
    mainDeliverOrderNo   int64 

    // 同意理由
    operateDesc   string 

}

func NewTmallChannelTradeDeliverorderAgreeRequest() *TmallChannelTradeDeliverorderAgreeRequest{
    return &TmallChannelTradeDeliverorderAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeDeliverorderAgreeRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.agree"
}

func (r TmallChannelTradeDeliverorderAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeDeliverorderAgreeRequest) SetMainDeliverOrderNo(mainDeliverOrderNo int64) error {
    r.mainDeliverOrderNo = mainDeliverOrderNo
    r.Set("main_deliver_order_no", mainDeliverOrderNo)
    return nil
}

func (r TmallChannelTradeDeliverorderAgreeRequest) GetMainDeliverOrderNo() int64 {
    return r.mainDeliverOrderNo
}

func (r *TmallChannelTradeDeliverorderAgreeRequest) SetOperateDesc(operateDesc string) error {
    r.operateDesc = operateDesc
    r.Set("operate_desc", operateDesc)
    return nil
}

func (r TmallChannelTradeDeliverorderAgreeRequest) GetOperateDesc() string {
    return r.operateDesc
}

