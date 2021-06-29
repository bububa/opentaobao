package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装业务查询物流公司api API请求
taobao.wlb.order.jz.query

家装业务查询物流公司api
*/
type TaobaoWlbOrderJzQueryRequest struct {
    model.Params
    // 交易id
    _tid   int64
    // 家装收货人信息
    _jzReceiverTo   *JzReceiverTo
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    _senderId   int64
    // 家装安装服务收货人信息
    _insJzReceiverTO   *JzReceiverTo
}

// 初始化TaobaoWlbOrderJzQueryRequest对象
func NewTaobaoWlbOrderJzQueryRequest() *TaobaoWlbOrderJzQueryRequest{
    return &TaobaoWlbOrderJzQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.order.jz.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 交易id
func (r *TaobaoWlbOrderJzQueryRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoWlbOrderJzQueryRequest) GetTid() int64 {
    return r._tid
}
// JzReceiverTo Setter
// 家装收货人信息
func (r *TaobaoWlbOrderJzQueryRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverTo) error {
    r._jzReceiverTo = _jzReceiverTo
    r.Set("jz_receiver_to", _jzReceiverTo)
    return nil
}

// JzReceiverTo Getter
func (r TaobaoWlbOrderJzQueryRequest) GetJzReceiverTo() *JzReceiverTo {
    return r._jzReceiverTo
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoWlbOrderJzQueryRequest) SetSenderId(_senderId int64) error {
    r._senderId = _senderId
    r.Set("sender_id", _senderId)
    return nil
}

// SenderId Getter
func (r TaobaoWlbOrderJzQueryRequest) GetSenderId() int64 {
    return r._senderId
}
// InsJzReceiverTO Setter
// 家装安装服务收货人信息
func (r *TaobaoWlbOrderJzQueryRequest) SetInsJzReceiverTO(_insJzReceiverTO *JzReceiverTo) error {
    r._insJzReceiverTO = _insJzReceiverTO
    r.Set("ins_jz_receiver_t_o", _insJzReceiverTO)
    return nil
}

// InsJzReceiverTO Getter
func (r TaobaoWlbOrderJzQueryRequest) GetInsJzReceiverTO() *JzReceiverTo {
    return r._insJzReceiverTO
}
