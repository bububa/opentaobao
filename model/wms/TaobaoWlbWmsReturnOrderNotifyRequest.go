package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退货通知 APIRequest
taobao.wlb.wms.return.order.notify

销售退货通知
*/
type TaobaoWlbWmsReturnOrderNotifyRequest struct {
    model.Params

    // ERP单据编码
    orderCode   string 

    // 仓库编码
    storeCode   string 

    // 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    orderFlag   string 

    // 来源单据号，销售退货时填充原发货的LBX号
    prevOrderCode   string 

    // 快递公司编码
    tmsServiceCode   string 

    // 运单号
    tmsOrderCode   string 

    // 销退时请提供退货的原因
    returnReason   string 

    // 买家昵称
    buyerNick   string 

    // 发件人信息
    senderInfo   *Senderinfowlbwmsreturnordernotify 

    // 收件人信息
    receiverInfo   *Receiverinfowlbwmsreturnordernotify 

    // 商品信息列表
    orderItemList   []Orderitemlistwlbwmsreturnordernotify 

    // 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    extendFields   string 

    // 备注
    remark   string 

    // 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
    orderSource   string 

    // 订单类型 501 销退入库
    orderType   string 

    // 货主ID
    ownerUserId   string 

    // ERP订单创建时间
    orderCreateTime   string 

    // 快递公司名称
    tmsServiceName   string 

}

func NewTaobaoWlbWmsReturnOrderNotifyRequest() *TaobaoWlbWmsReturnOrderNotifyRequest{
    return &TaobaoWlbWmsReturnOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.return.order.notify"
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderFlag(orderFlag string) error {
    r.orderFlag = orderFlag
    r.Set("order_flag", orderFlag)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderFlag() string {
    return r.orderFlag
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsServiceCode(tmsServiceCode string) error {
    r.tmsServiceCode = tmsServiceCode
    r.Set("tms_service_code", tmsServiceCode)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsServiceCode() string {
    return r.tmsServiceCode
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsOrderCode(tmsOrderCode string) error {
    r.tmsOrderCode = tmsOrderCode
    r.Set("tms_order_code", tmsOrderCode)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsOrderCode() string {
    return r.tmsOrderCode
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetReturnReason(returnReason string) error {
    r.returnReason = returnReason
    r.Set("return_reason", returnReason)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetReturnReason() string {
    return r.returnReason
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetSenderInfo(senderInfo *Senderinfowlbwmsreturnordernotify) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetSenderInfo() *Senderinfowlbwmsreturnordernotify {
    return r.senderInfo
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetReceiverInfo(receiverInfo *Receiverinfowlbwmsreturnordernotify) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetReceiverInfo() *Receiverinfowlbwmsreturnordernotify {
    return r.receiverInfo
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderItemList(orderItemList []Orderitemlistwlbwmsreturnordernotify) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsreturnordernotify {
    return r.orderItemList
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetExtendFields() string {
    return r.extendFields
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderSource(orderSource string) error {
    r.orderSource = orderSource
    r.Set("order_source", orderSource)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderSource() string {
    return r.orderSource
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderType() string {
    return r.orderType
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOwnerUserId(ownerUserId string) error {
    r.ownerUserId = ownerUserId
    r.Set("owner_user_id", ownerUserId)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOwnerUserId() string {
    return r.ownerUserId
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetOrderCreateTime(orderCreateTime string) error {
    r.orderCreateTime = orderCreateTime
    r.Set("order_create_time", orderCreateTime)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetOrderCreateTime() string {
    return r.orderCreateTime
}

func (r *TaobaoWlbWmsReturnOrderNotifyRequest) SetTmsServiceName(tmsServiceName string) error {
    r.tmsServiceName = tmsServiceName
    r.Set("tms_service_name", tmsServiceName)
    return nil
}

func (r TaobaoWlbWmsReturnOrderNotifyRequest) GetTmsServiceName() string {
    return r.tmsServiceName
}

