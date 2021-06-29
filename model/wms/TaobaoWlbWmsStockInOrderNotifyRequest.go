package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库通知单 API请求
taobao.wlb.wms.stock.in.order.notify

入库通知单
*/
type TaobaoWlbWmsStockInOrderNotifyRequest struct {
    model.Params
    // 仓库编码
    _storeCode   string
    // 入库单据编码
    _orderCode   string
    // 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
    _orderType   int64
    // 可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
    _inboundTypeDesc   string
    // 订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
    _orderFlag   string
    // 单据创建时间
    _orderCreateTime   string
    // 供应商编码，往来单位编码
    _supplierCode   string
    // 供应商名称 ，往来单位名称
    _supplierName   string
    // 配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
    _tmsServiceCode   string
    // 快递公司名称
    _tmsServiceName   string
    // 运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
    _tmsOrderCode   string
    // 来源单据号，如销售退货时填充原销售订单号
    _prevOrderCode   string
    // 销退时请提供退货的原因
    _returnReason   string
    // 预期送达开始时间
    _expectStartTime   string
    // 预期送达结束时间
    _expectEndTime   string
    // 系统自动生成
    _receiverInfo   *Receiverinfowlbwmsstockinordernotifywl
    // 系统自动生成
    _senderInfo   *Senderinfowlbwmsstockinordernotifywl
    // 系统自动生成
    _orderItemList   []Orderitemlistwlbwmsstockinordernotifywl
    // 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
    _extendFields   string
    // 备注，销退入库订单需要留言备注填充到此字段
    _remark   string
}

// 初始化TaobaoWlbWmsStockInOrderNotifyRequest对象
func NewTaobaoWlbWmsStockInOrderNotifyRequest() *TaobaoWlbWmsStockInOrderNotifyRequest{
    return &TaobaoWlbWmsStockInOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.in.order.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetStoreCode() string {
    return r._storeCode
}
// OrderCode Setter
// 入库单据编码
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderCode() string {
    return r._orderCode
}
// OrderType Setter
// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderType() int64 {
    return r._orderType
}
// InboundTypeDesc Setter
// 可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetInboundTypeDesc(_inboundTypeDesc string) error {
    r._inboundTypeDesc = _inboundTypeDesc
    r.Set("inbound_type_desc", _inboundTypeDesc)
    return nil
}

// InboundTypeDesc Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetInboundTypeDesc() string {
    return r._inboundTypeDesc
}
// OrderFlag Setter
// 订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderFlag(_orderFlag string) error {
    r._orderFlag = _orderFlag
    r.Set("order_flag", _orderFlag)
    return nil
}

// OrderFlag Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderFlag() string {
    return r._orderFlag
}
// OrderCreateTime Setter
// 单据创建时间
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderCreateTime(_orderCreateTime string) error {
    r._orderCreateTime = _orderCreateTime
    r.Set("order_create_time", _orderCreateTime)
    return nil
}

// OrderCreateTime Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderCreateTime() string {
    return r._orderCreateTime
}
// SupplierCode Setter
// 供应商编码，往来单位编码
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetSupplierCode(_supplierCode string) error {
    r._supplierCode = _supplierCode
    r.Set("supplier_code", _supplierCode)
    return nil
}

// SupplierCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetSupplierCode() string {
    return r._supplierCode
}
// SupplierName Setter
// 供应商名称 ，往来单位名称
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetSupplierName(_supplierName string) error {
    r._supplierName = _supplierName
    r.Set("supplier_name", _supplierName)
    return nil
}

// SupplierName Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetSupplierName() string {
    return r._supplierName
}
// TmsServiceCode Setter
// 配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsServiceCode(_tmsServiceCode string) error {
    r._tmsServiceCode = _tmsServiceCode
    r.Set("tms_service_code", _tmsServiceCode)
    return nil
}

// TmsServiceCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetTmsServiceCode() string {
    return r._tmsServiceCode
}
// TmsServiceName Setter
// 快递公司名称
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsServiceName(_tmsServiceName string) error {
    r._tmsServiceName = _tmsServiceName
    r.Set("tms_service_name", _tmsServiceName)
    return nil
}

// TmsServiceName Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetTmsServiceName() string {
    return r._tmsServiceName
}
// TmsOrderCode Setter
// 运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsOrderCode(_tmsOrderCode string) error {
    r._tmsOrderCode = _tmsOrderCode
    r.Set("tms_order_code", _tmsOrderCode)
    return nil
}

// TmsOrderCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetTmsOrderCode() string {
    return r._tmsOrderCode
}
// PrevOrderCode Setter
// 来源单据号，如销售退货时填充原销售订单号
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetPrevOrderCode(_prevOrderCode string) error {
    r._prevOrderCode = _prevOrderCode
    r.Set("prev_order_code", _prevOrderCode)
    return nil
}

// PrevOrderCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetPrevOrderCode() string {
    return r._prevOrderCode
}
// ReturnReason Setter
// 销退时请提供退货的原因
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetReturnReason(_returnReason string) error {
    r._returnReason = _returnReason
    r.Set("return_reason", _returnReason)
    return nil
}

// ReturnReason Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetReturnReason() string {
    return r._returnReason
}
// ExpectStartTime Setter
// 预期送达开始时间
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetExpectStartTime(_expectStartTime string) error {
    r._expectStartTime = _expectStartTime
    r.Set("expect_start_time", _expectStartTime)
    return nil
}

// ExpectStartTime Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetExpectStartTime() string {
    return r._expectStartTime
}
// ExpectEndTime Setter
// 预期送达结束时间
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetExpectEndTime(_expectEndTime string) error {
    r._expectEndTime = _expectEndTime
    r.Set("expect_end_time", _expectEndTime)
    return nil
}

// ExpectEndTime Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetExpectEndTime() string {
    return r._expectEndTime
}
// ReceiverInfo Setter
// 系统自动生成
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetReceiverInfo(_receiverInfo *Receiverinfowlbwmsstockinordernotifywl) error {
    r._receiverInfo = _receiverInfo
    r.Set("receiver_info", _receiverInfo)
    return nil
}

// ReceiverInfo Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetReceiverInfo() *Receiverinfowlbwmsstockinordernotifywl {
    return r._receiverInfo
}
// SenderInfo Setter
// 系统自动生成
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetSenderInfo(_senderInfo *Senderinfowlbwmsstockinordernotifywl) error {
    r._senderInfo = _senderInfo
    r.Set("sender_info", _senderInfo)
    return nil
}

// SenderInfo Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetSenderInfo() *Senderinfowlbwmsstockinordernotifywl {
    return r._senderInfo
}
// OrderItemList Setter
// 系统自动生成
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsstockinordernotifywl) error {
    r._orderItemList = _orderItemList
    r.Set("order_item_list", _orderItemList)
    return nil
}

// OrderItemList Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetOrderItemList() []Orderitemlistwlbwmsstockinordernotifywl {
    return r._orderItemList
}
// ExtendFields Setter
// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetExtendFields(_extendFields string) error {
    r._extendFields = _extendFields
    r.Set("extend_fields", _extendFields)
    return nil
}

// ExtendFields Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetExtendFields() string {
    return r._extendFields
}
// Remark Setter
// 备注，销退入库订单需要留言备注填充到此字段
func (r *TaobaoWlbWmsStockInOrderNotifyRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbWmsStockInOrderNotifyRequest) GetRemark() string {
    return r._remark
}
