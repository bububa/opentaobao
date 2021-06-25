package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库通知单 APIRequest
taobao.wlb.wms.stock.in.order.notify

入库通知单
*/
type TaobaoWlbWmsStockInOrderNotifyRequest struct {
    model.Params

    // 仓库编码
    storeCode   string 

    // 入库单据编码
    orderCode   string 

    // 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
    orderType   int64 

    // 可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
    inboundTypeDesc   string 

    // 订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
    orderFlag   string 

    // 单据创建时间
    orderCreateTime   string 

    // 供应商编码，往来单位编码
    supplierCode   string 

    // 供应商名称 ，往来单位名称
    supplierName   string 

    // 配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
    tmsServiceCode   string 

    // 快递公司名称
    tmsServiceName   string 

    // 运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
    tmsOrderCode   string 

    // 来源单据号，如销售退货时填充原销售订单号
    prevOrderCode   string 

    // 销退时请提供退货的原因
    returnReason   string 

    // 预期送达开始时间
    expectStartTime   string 

    // 预期送达结束时间
    expectEndTime   string 

    // 系统自动生成
    receiverInfo   *Receiverinfowlbwmsstockinordernotifywl 

    // 系统自动生成
    senderInfo   *Senderinfowlbwmsstockinordernotifywl 

    // 系统自动生成
    orderItemList   []Orderitemlistwlbwmsstockinordernotifywl 

    // 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    extendFields   string 

    // 备注，销退入库订单需要留言备注填充到此字段
    remark   string 

}

func NewTaobaoWlbWmsStockInOrderNotifyRequest() *TaobaoWlbWmsStockInOrderNotifyRequest{
    return &TaobaoWlbWmsStockInOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.in.order.notify"
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetInboundTypeDesc(inboundTypeDesc string) error {
    r.inboundTypeDesc = inboundTypeDesc
    r.Set("inbound_type_desc", inboundTypeDesc)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetInboundTypeDesc() string {
    return r.inboundTypeDesc
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderFlag(orderFlag string) error {
    r.orderFlag = orderFlag
    r.Set("order_flag", orderFlag)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderFlag() string {
    return r.orderFlag
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderCreateTime(orderCreateTime string) error {
    r.orderCreateTime = orderCreateTime
    r.Set("order_create_time", orderCreateTime)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderCreateTime() string {
    return r.orderCreateTime
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetSupplierCode(supplierCode string) error {
    r.supplierCode = supplierCode
    r.Set("supplier_code", supplierCode)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetSupplierCode() string {
    return r.supplierCode
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetSupplierName(supplierName string) error {
    r.supplierName = supplierName
    r.Set("supplier_name", supplierName)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetSupplierName() string {
    return r.supplierName
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsServiceCode(tmsServiceCode string) error {
    r.tmsServiceCode = tmsServiceCode
    r.Set("tms_service_code", tmsServiceCode)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetTmsServiceCode() string {
    return r.tmsServiceCode
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsServiceName(tmsServiceName string) error {
    r.tmsServiceName = tmsServiceName
    r.Set("tms_service_name", tmsServiceName)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetTmsServiceName() string {
    return r.tmsServiceName
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsOrderCode(tmsOrderCode string) error {
    r.tmsOrderCode = tmsOrderCode
    r.Set("tms_order_code", tmsOrderCode)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetTmsOrderCode() string {
    return r.tmsOrderCode
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetReturnReason(returnReason string) error {
    r.returnReason = returnReason
    r.Set("return_reason", returnReason)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetReturnReason() string {
    return r.returnReason
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetExpectStartTime(expectStartTime string) error {
    r.expectStartTime = expectStartTime
    r.Set("expect_start_time", expectStartTime)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetExpectStartTime() string {
    return r.expectStartTime
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetExpectEndTime(expectEndTime string) error {
    r.expectEndTime = expectEndTime
    r.Set("expect_end_time", expectEndTime)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetExpectEndTime() string {
    return r.expectEndTime
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetReceiverInfo(receiverInfo *Receiverinfowlbwmsstockinordernotifywl) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetReceiverInfo() *Receiverinfowlbwmsstockinordernotifywl {
    return r.receiverInfo
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetSenderInfo(senderInfo *Senderinfowlbwmsstockinordernotifywl) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetSenderInfo() *Senderinfowlbwmsstockinordernotifywl {
    return r.senderInfo
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderItemList(orderItemList []Orderitemlistwlbwmsstockinordernotifywl) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsstockinordernotifywl {
    return r.orderItemList
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetExtendFields(extendFields string) error {
    r.extendFields = extendFields
    r.Set("extend_fields", extendFields)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetExtendFields() string {
    return r.extendFields
}

func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetRemark() string {
    return r.remark
}

