package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单通知 APIRequest
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

func NewTaobaoWlbWmsStockOutOrderNotifyRequest() *TaobaoWlbWmsStockOutOrderNotifyRequest{
    return &TaobaoWlbWmsStockOutOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.out.order.notify"
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOutboundTypeDesc(outboundTypeDesc string) error {
    r.outboundTypeDesc = outboundTypeDesc
    r.Set("outbound_type_desc", outboundTypeDesc)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOutboundTypeDesc() string {
    return r.outboundTypeDesc
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderCreateTime(orderCreateTime string) error {
    r.orderCreateTime = orderCreateTime
    r.Set("order_create_time", orderCreateTime)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderCreateTime() string {
    return r.orderCreateTime
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetSendTime(sendTime string) error {
    r.sendTime = sendTime
    r.Set("send_time", sendTime)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetSendTime() string {
    return r.sendTime
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetReceiverInfo(receiverInfo *Receiverwlbwmsstockoutordernotify) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetReceiverInfo() *Receiverwlbwmsstockoutordernotify {
    return r.receiverInfo
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetSenderInfo(senderInfo *Senderwlbwmsstockoutordernotify) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetSenderInfo() *Senderwlbwmsstockoutordernotify {
    return r.senderInfo
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetTransportMode(transportMode string) error {
    r.transportMode = transportMode
    r.Set("transport_mode", transportMode)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetTransportMode() string {
    return r.transportMode
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetCarriersName(carriersName string) error {
    r.carriersName = carriersName
    r.Set("carriers_name", carriersName)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetCarriersName() string {
    return r.carriersName
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickName(pickName string) error {
    r.pickName = pickName
    r.Set("pick_name", pickName)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickName() string {
    return r.pickName
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickCall(pickCall string) error {
    r.pickCall = pickCall
    r.Set("pick_call", pickCall)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickCall() string {
    return r.pickCall
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPickId(pickId string) error {
    r.pickId = pickId
    r.Set("pick_id", pickId)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPickId() string {
    return r.pickId
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetCarNo(carNo string) error {
    r.carNo = carNo
    r.Set("car_no", carNo)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetCarNo() string {
    return r.carNo
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetOrderItemList(orderItemList []Orderitemlistwlbwmsstockoutordernotify) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsstockoutordernotify {
    return r.orderItemList
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}

func (r *TaobaoWlbWmsStockOutOrderNotifyRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

func (r TaobaoWlbWmsStockOutOrderNotifyRequest) GetExtendFields() string {
    return r.extendFields
}

