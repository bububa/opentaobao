package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单通知 API请求
taobao.wlb.wms.stock.out.order.notify

出库单通知
*/
type TaobaoWlbWmsStockOutOrderNotifyRequest struct {
    model.Params
    // 仓储编码
    storeCode   string
    // ERP单据ID
    orderCode   string
    // 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
    orderType   int64
    // ERP可选择性文本透传至WMS
    outboundTypeDesc   string
    // 订单创建时间
    orderCreateTime   string
    // 要求出库日期
    sendTime   string
    // 收件人信息
    receiverInfo   *Receiverwlbwmsstockoutordernotify
    // 发货方信息
    senderInfo   *Senderwlbwmsstockoutordernotify
    // 出库方式
    transportMode   string
    // 承运商名称
    carriersName   string
    // 取件人姓名
    pickName   string
    // 取件人电话
    pickCall   string
    // 取件人身份证ID
    pickId   string
    // 车牌号
    carNo   string
    // 订单商品信息列表
    orderItemList   []Orderitemlistwlbwmsstockoutordernotify
    // 备注
    remark   string
    // 前物流订单号，如退货入库单可能会用到
    prevOrderCode   string
    // 拓展属性
    extendFields   string
}

// 初始化TaobaoWlbWmsStockOutOrderNotifyRequest对象
func NewTaobaoWlbWmsStockOutOrderNotifyRequest() *TaobaoWlbWmsStockOutOrderNotifyRequest{
    return &TaobaoWlbWmsStockOutOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.out.order.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓储编码
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetStoreCode() string {
    return r.storeCode
}
// OrderCode Setter
// ERP单据ID
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderCode() string {
    return r.orderCode
}
// OrderType Setter
// 单据类型 301 调拨出库单、901普通出库单、903 其他出库单 305 B2B出库
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderType() int64 {
    return r.orderType
}
// OutboundTypeDesc Setter
// ERP可选择性文本透传至WMS
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOutboundTypeDesc(outboundTypeDesc string) error {
    r.outboundTypeDesc = outboundTypeDesc
    r.Set("outbound_type_desc", outboundTypeDesc)
    return nil
}

// OutboundTypeDesc Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOutboundTypeDesc() string {
    return r.outboundTypeDesc
}
// OrderCreateTime Setter
// 订单创建时间
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderCreateTime(orderCreateTime string) error {
    r.orderCreateTime = orderCreateTime
    r.Set("order_create_time", orderCreateTime)
    return nil
}

// OrderCreateTime Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderCreateTime() string {
    return r.orderCreateTime
}
// SendTime Setter
// 要求出库日期
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetSendTime(sendTime string) error {
    r.sendTime = sendTime
    r.Set("send_time", sendTime)
    return nil
}

// SendTime Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetSendTime() string {
    return r.sendTime
}
// ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetReceiverInfo(receiverInfo *Receiverwlbwmsstockoutordernotify) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

// ReceiverInfo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetReceiverInfo() *Receiverwlbwmsstockoutordernotify {
    return r.receiverInfo
}
// SenderInfo Setter
// 发货方信息
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetSenderInfo(senderInfo *Senderwlbwmsstockoutordernotify) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

// SenderInfo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetSenderInfo() *Senderwlbwmsstockoutordernotify {
    return r.senderInfo
}
// TransportMode Setter
// 出库方式
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetTransportMode(transportMode string) error {
    r.transportMode = transportMode
    r.Set("transport_mode", transportMode)
    return nil
}

// TransportMode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetTransportMode() string {
    return r.transportMode
}
// CarriersName Setter
// 承运商名称
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetCarriersName(carriersName string) error {
    r.carriersName = carriersName
    r.Set("carriers_name", carriersName)
    return nil
}

// CarriersName Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetCarriersName() string {
    return r.carriersName
}
// PickName Setter
// 取件人姓名
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickName(pickName string) error {
    r.pickName = pickName
    r.Set("pick_name", pickName)
    return nil
}

// PickName Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickName() string {
    return r.pickName
}
// PickCall Setter
// 取件人电话
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickCall(pickCall string) error {
    r.pickCall = pickCall
    r.Set("pick_call", pickCall)
    return nil
}

// PickCall Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickCall() string {
    return r.pickCall
}
// PickId Setter
// 取件人身份证ID
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickId(pickId string) error {
    r.pickId = pickId
    r.Set("pick_id", pickId)
    return nil
}

// PickId Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickId() string {
    return r.pickId
}
// CarNo Setter
// 车牌号
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetCarNo(carNo string) error {
    r.carNo = carNo
    r.Set("car_no", carNo)
    return nil
}

// CarNo Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetCarNo() string {
    return r.carNo
}
// OrderItemList Setter
// 订单商品信息列表
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderItemList(orderItemList []Orderitemlistwlbwmsstockoutordernotify) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

// OrderItemList Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsstockoutordernotify {
    return r.orderItemList
}
// Remark Setter
// 备注
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetRemark() string {
    return r.remark
}
// PrevOrderCode Setter
// 前物流订单号，如退货入库单可能会用到
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

// PrevOrderCode Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}
// ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetExtendFields() string {
    return r.extendFields
}
