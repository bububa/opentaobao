package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装业务查询物流公司api APIRequest
taobao.wlb.order.jz.query

家装业务查询物流公司api
*/
type TaobaoWlbOrderJzQueryRequest struct {
    model.Params

    // 交易id
    tid   int64 

    // 家装收货人信息
    jzReceiverTo   *JzReceiverTo 

    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    senderId   int64 

    // 家装安装服务收货人信息
    insJzReceiverTO   *JzReceiverTo 

}

func NewTaobaoWlbOrderJzQueryRequest() *TaobaoWlbOrderJzQueryRequest{
    return &TaobaoWlbOrderJzQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderJzQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.order.jz.query"
}

func (r TaobaoWlbOrderJzQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderJzQueryRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoWlbOrderJzQueryRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoWlbOrderJzQueryRequest) SetJzReceiverTo(jzReceiverTo *JzReceiverTo) error {
    r.jzReceiverTo = jzReceiverTo
    r.Set("jz_receiver_to", jzReceiverTo)
    return nil
}

func (r TaobaoWlbOrderJzQueryRequest) GetJzReceiverTo() *JzReceiverTo {
    return r.jzReceiverTo
}

func (r *TaobaoWlbOrderJzQueryRequest) SetSenderId(senderId int64) error {
    r.senderId = senderId
    r.Set("sender_id", senderId)
    return nil
}

func (r TaobaoWlbOrderJzQueryRequest) GetSenderId() int64 {
    return r.senderId
}

func (r *TaobaoWlbOrderJzQueryRequest) SetInsJzReceiverTO(insJzReceiverTO *JzReceiverTo) error {
    r.insJzReceiverTO = insJzReceiverTO
    r.Set("ins_jz_receiver_t_o", insJzReceiverTO)
    return nil
}

func (r TaobaoWlbOrderJzQueryRequest) GetInsJzReceiverTO() *JzReceiverTo {
    return r.insJzReceiverTO
}

