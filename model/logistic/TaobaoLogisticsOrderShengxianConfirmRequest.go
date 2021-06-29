package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝生鲜冷链的发货 API请求
taobao.logistics.order.shengxian.confirm

优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
*/
type TaobaoLogisticsOrderShengxianConfirmRequest struct {
    model.Params
    // 淘宝交易ID
    _tid   int64
    // 运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入；
    _outSid   string
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货
    _senderId   int64
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br>
    _cancelId   int64
    // 商家的IP地址
    _sellerIp   string
    // 物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准
    _logisId   int64
    // 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货
    _serviceCode   string
    // 1：冷链。0：常温
    _deliveryType   int64
}

// 初始化TaobaoLogisticsOrderShengxianConfirmRequest对象
func NewTaobaoLogisticsOrderShengxianConfirmRequest() *TaobaoLogisticsOrderShengxianConfirmRequest{
    return &TaobaoLogisticsOrderShengxianConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetApiMethodName() string {
    return "taobao.logistics.order.shengxian.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetTid() int64 {
    return r._tid
}
// OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetOutSid(_outSid string) error {
    r._outSid = _outSid
    r.Set("out_sid", _outSid)
    return nil
}

// OutSid Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetOutSid() string {
    return r._outSid
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetSenderId(_senderId int64) error {
    r._senderId = _senderId
    r.Set("sender_id", _senderId)
    return nil
}

// SenderId Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetSenderId() int64 {
    return r._senderId
}
// CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br>
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetCancelId(_cancelId int64) error {
    r._cancelId = _cancelId
    r.Set("cancel_id", _cancelId)
    return nil
}

// CancelId Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetCancelId() int64 {
    return r._cancelId
}
// SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetSellerIp(_sellerIp string) error {
    r._sellerIp = _sellerIp
    r.Set("seller_ip", _sellerIp)
    return nil
}

// SellerIp Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetSellerIp() string {
    return r._sellerIp
}
// LogisId Setter
// 物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetLogisId(_logisId int64) error {
    r._logisId = _logisId
    r.Set("logis_id", _logisId)
    return nil
}

// LogisId Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetLogisId() int64 {
    return r._logisId
}
// ServiceCode Setter
// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetServiceCode(_serviceCode string) error {
    r._serviceCode = _serviceCode
    r.Set("service_code", _serviceCode)
    return nil
}

// ServiceCode Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetServiceCode() string {
    return r._serviceCode
}
// DeliveryType Setter
// 1：冷链。0：常温
func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetDeliveryType(_deliveryType int64) error {
    r._deliveryType = _deliveryType
    r.Set("delivery_type", _deliveryType)
    return nil
}

// DeliveryType Getter
func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetDeliveryType() int64 {
    return r._deliveryType
}
