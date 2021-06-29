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
    _tid   int64
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    _senderId   int64
    // 家装收货人信息,如果为空,则取默认收货信息
    _jzReceiverTo   *JzReceiverTo
    // 发货参数
    _jzTopArgs   *JzTopArgs
    // 物流公司信息
    _lgTpDto   *Tpdto
    // 安装公司信息,需要安装时,才填写
    _insTpDto   *Tpdto
    // 安装收货人信息,如果为空,则取默认收货人信息
    _insReceiverTo   *JzReceiverTo
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
func (r *TaobaoWlbOrderJzConsignRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoWlbOrderJzConsignRequest) GetTid() int64 {
    return r._tid
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoWlbOrderJzConsignRequest) SetSenderId(_senderId int64) error {
    r._senderId = _senderId
    r.Set("sender_id", _senderId)
    return nil
}

// SenderId Getter
func (r TaobaoWlbOrderJzConsignRequest) GetSenderId() int64 {
    return r._senderId
}
// JzReceiverTo Setter
// 家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaoWlbOrderJzConsignRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverTo) error {
    r._jzReceiverTo = _jzReceiverTo
    r.Set("jz_receiver_to", _jzReceiverTo)
    return nil
}

// JzReceiverTo Getter
func (r TaobaoWlbOrderJzConsignRequest) GetJzReceiverTo() *JzReceiverTo {
    return r._jzReceiverTo
}
// JzTopArgs Setter
// 发货参数
func (r *TaobaoWlbOrderJzConsignRequest) SetJzTopArgs(_jzTopArgs *JzTopArgs) error {
    r._jzTopArgs = _jzTopArgs
    r.Set("jz_top_args", _jzTopArgs)
    return nil
}

// JzTopArgs Getter
func (r TaobaoWlbOrderJzConsignRequest) GetJzTopArgs() *JzTopArgs {
    return r._jzTopArgs
}
// LgTpDto Setter
// 物流公司信息
func (r *TaobaoWlbOrderJzConsignRequest) SetLgTpDto(_lgTpDto *Tpdto) error {
    r._lgTpDto = _lgTpDto
    r.Set("lg_tp_dto", _lgTpDto)
    return nil
}

// LgTpDto Getter
func (r TaobaoWlbOrderJzConsignRequest) GetLgTpDto() *Tpdto {
    return r._lgTpDto
}
// InsTpDto Setter
// 安装公司信息,需要安装时,才填写
func (r *TaobaoWlbOrderJzConsignRequest) SetInsTpDto(_insTpDto *Tpdto) error {
    r._insTpDto = _insTpDto
    r.Set("ins_tp_dto", _insTpDto)
    return nil
}

// InsTpDto Getter
func (r TaobaoWlbOrderJzConsignRequest) GetInsTpDto() *Tpdto {
    return r._insTpDto
}
// InsReceiverTo Setter
// 安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaoWlbOrderJzConsignRequest) SetInsReceiverTo(_insReceiverTo *JzReceiverTo) error {
    r._insReceiverTo = _insReceiverTo
    r.Set("ins_receiver_to", _insReceiverTo)
    return nil
}

// InsReceiverTo Getter
func (r TaobaoWlbOrderJzConsignRequest) GetInsReceiverTo() *JzReceiverTo {
    return r._insReceiverTo
}
