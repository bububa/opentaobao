package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
区域零售订单获取取件码 APIRequest
tmall.nr.order.logis.info

区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
*/
type TmallNrOrderLogisInfoRequest struct {
    model.Params

    // 卖家ID
    sellerId   int64 

    // 主订单号
    mainOrderIds   []int64 

    // 来源标识
    channel   string 

}

func NewTmallNrOrderLogisInfoRequest() *TmallNrOrderLogisInfoRequest{
    return &TmallNrOrderLogisInfoRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrOrderLogisInfoRequest) GetApiMethodName() string {
    return "tmall.nr.order.logis.info"
}

func (r TmallNrOrderLogisInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrOrderLogisInfoRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TmallNrOrderLogisInfoRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *TmallNrOrderLogisInfoRequest) SetMainOrderIds(mainOrderIds []int64) error {
    r.mainOrderIds = mainOrderIds
    r.Set("main_order_ids", mainOrderIds)
    return nil
}

func (r TmallNrOrderLogisInfoRequest) GetMainOrderIds() []int64 {
    return r.mainOrderIds
}

func (r *TmallNrOrderLogisInfoRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TmallNrOrderLogisInfoRequest) GetChannel() string {
    return r.channel
}

