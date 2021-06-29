package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装发货接口 API请求
taobao.wlb.order.jz.consign

家装类订单使用该接口发货
*/
type TaobaoWlbOrderJzConsignRequest struct {
    model.Params
    // 交易号
    tid   int64
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    senderId   int64
    // 家装收货人信息,如果为空,则取默认收货信息
    jzReceiverTo   *JzReceiverTo
    // 发货参数
    jzTopArgs   *JzTopArgs
    // 物流公司信息
    lgTpDto   *Tpdto
    // 安装公司信息,需要安装时,才填写
    insTpDto   *Tpdto
    // 安装收货人信息,如果为空,则取默认收货人信息
    insReceiverTo   *JzReceiverTo
}

// 初始化TaobaoWlbOrderJzConsignRequest对象
func NewTaobaoWlbOrderJzConsignRequest() *TaobaoWlbOrderJzConsignRequest{
    return &TaobaoWlbOrderJzConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzConsignRequest) GetApiMethodName() string {
    return "taobao.wlb.order.jz.consign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易号
func (r *TaobaoWlbOrderJzConsignRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoWlbOrderJzConsignRequest) GetTid() int64 {
    return r.tid
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoWlbOrderJzConsignRequest) SetSenderId(senderId int64) error {
    r.senderId = senderId
    r.Set("sender_id", senderId)
    return nil
}

// SenderId Getter
func (r TaobaoWlbOrderJzConsignRequest) GetSenderId() int64 {
    return r.senderId
}
// JzReceiverTo Setter
// 家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaoWlbOrderJzConsignRequest) SetJzReceiverTo(jzReceiverTo *JzReceiverTo) error {
    r.jzReceiverTo = jzReceiverTo
    r.Set("jz_receiver_to", jzReceiverTo)
    return nil
}

// JzReceiverTo Getter
func (r TaobaoWlbOrderJzConsignRequest) GetJzReceiverTo() *JzReceiverTo {
    return r.jzReceiverTo
}
// JzTopArgs Setter
// 发货参数
func (r *TaobaoWlbOrderJzConsignRequest) SetJzTopArgs(jzTopArgs *JzTopArgs) error {
    r.jzTopArgs = jzTopArgs
    r.Set("jz_top_args", jzTopArgs)
    return nil
}

// JzTopArgs Getter
func (r TaobaoWlbOrderJzConsignRequest) GetJzTopArgs() *JzTopArgs {
    return r.jzTopArgs
}
// LgTpDto Setter
// 物流公司信息
func (r *TaobaoWlbOrderJzConsignRequest) SetLgTpDto(lgTpDto *Tpdto) error {
    r.lgTpDto = lgTpDto
    r.Set("lg_tp_dto", lgTpDto)
    return nil
}

// LgTpDto Getter
func (r TaobaoWlbOrderJzConsignRequest) GetLgTpDto() *Tpdto {
    return r.lgTpDto
}
// InsTpDto Setter
// 安装公司信息,需要安装时,才填写
func (r *TaobaoWlbOrderJzConsignRequest) SetInsTpDto(insTpDto *Tpdto) error {
    r.insTpDto = insTpDto
    r.Set("ins_tp_dto", insTpDto)
    return nil
}

// InsTpDto Getter
func (r TaobaoWlbOrderJzConsignRequest) GetInsTpDto() *Tpdto {
    return r.insTpDto
}
// InsReceiverTo Setter
// 安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaoWlbOrderJzConsignRequest) SetInsReceiverTo(insReceiverTo *JzReceiverTo) error {
    r.insReceiverTo = insReceiverTo
    r.Set("ins_receiver_to", insReceiverTo)
    return nil
}

// InsReceiverTo Getter
func (r TaobaoWlbOrderJzConsignRequest) GetInsReceiverTo() *JzReceiverTo {
    return r.insReceiverTo
}
