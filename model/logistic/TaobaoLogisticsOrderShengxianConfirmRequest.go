package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝生鲜冷链的发货 APIRequest
taobao.logistics.order.shengxian.confirm

优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
*/
type TaobaoLogisticsOrderShengxianConfirmRequest struct {
    model.Params

    // 淘宝交易ID
    tid   int64 

    // 运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入；
    outSid   string 

    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货
    senderId   int64 

    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br>
    cancelId   int64 

    // 商家的IP地址
    sellerIp   string 

    // 物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准
    logisId   int64 

    // 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货
    serviceCode   string 

    // 1：冷链。0：常温
    deliveryType   int64 

}

func NewTaobaoLogisticsOrderShengxianConfirmRequest() *TaobaoLogisticsOrderShengxianConfirmRequest{
    return &TaobaoLogisticsOrderShengxianConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetApiMethodName() string {
    return "taobao.logistics.order.shengxian.confirm"
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetOutSid(outSid string) error {
    r.outSid = outSid
    r.Set("out_sid", outSid)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetOutSid() string {
    return r.outSid
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetSenderId(senderId int64) error {
    r.senderId = senderId
    r.Set("sender_id", senderId)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetSenderId() int64 {
    return r.senderId
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetCancelId(cancelId int64) error {
    r.cancelId = cancelId
    r.Set("cancel_id", cancelId)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetCancelId() int64 {
    return r.cancelId
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetSellerIp(sellerIp string) error {
    r.sellerIp = sellerIp
    r.Set("seller_ip", sellerIp)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetSellerIp() string {
    return r.sellerIp
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetLogisId(logisId int64) error {
    r.logisId = logisId
    r.Set("logis_id", logisId)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetLogisId() int64 {
    return r.logisId
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetServiceCode() string {
    return r.serviceCode
}

func (r *TaobaoLogisticsOrderShengxianConfirmRequest) SetDeliveryType(deliveryType int64) error {
    r.deliveryType = deliveryType
    r.Set("delivery_type", deliveryType)
    return nil
}

func (r TaobaoLogisticsOrderShengxianConfirmRequest) GetDeliveryType() int64 {
    return r.deliveryType
}

