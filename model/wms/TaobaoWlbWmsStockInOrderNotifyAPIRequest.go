package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsStockInOrderNotifyAPIRequest
入库通知单 API请求
taobao.wlb.wms.stock.in.order.notify

入库通知单 */
type TaobaoWlbWmsStockInOrderNotifyAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
	// 入库单据编码
	_orderCode string
	// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
	_orderType int64
	// 可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
	_inboundTypeDesc string
	// 订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
	_orderFlag string
	// 单据创建时间
	_orderCreateTime string
	// 供应商编码，往来单位编码
	_supplierCode string
	// 供应商名称 ，往来单位名称
	_supplierName string
	// 配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
	_tmsServiceCode string
	// 快递公司名称
	_tmsServiceName string
	// 运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
	_tmsOrderCode string
	// 来源单据号，如销售退货时填充原销售订单号
	_prevOrderCode string
	// 销退时请提供退货的原因
	_returnReason string
	// 预期送达开始时间
	_expectStartTime string
	// 预期送达结束时间
	_expectEndTime string
	// 系统自动生成
	_receiverInfo *Receiverinfowlbwmsstockinordernotifywl
	// 系统自动生成
	_senderInfo *Senderinfowlbwmsstockinordernotifywl
	// 系统自动生成
	_orderItemList []Orderitemlistwlbwmsstockinordernotifywl
	// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
	_extendFields string
	// 备注，销退入库订单需要留言备注填充到此字段
	_remark string
}

// NewTaobaoWlbWmsStockInOrderNotifyRequest 初始化TaobaoWlbWmsStockInOrderNotifyAPIRequest对象
func NewTaobaoWlbWmsStockInOrderNotifyRequest() *TaobaoWlbWmsStockInOrderNotifyAPIRequest {
	return &TaobaoWlbWmsStockInOrderNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.stock.in.order.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// Get StoreCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// Set is OrderCode Setter
// 入库单据编码
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// Get OrderCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// Set is OrderType Setter
// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// Get OrderType Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// Set is InboundTypeDesc Setter
// 可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetInboundTypeDesc(_inboundTypeDesc string) error {
	r._inboundTypeDesc = _inboundTypeDesc
	r.Set("inbound_type_desc", _inboundTypeDesc)
	return nil
}

// Get InboundTypeDesc Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetInboundTypeDesc() string {
	return r._inboundTypeDesc
}

// Set is OrderFlag Setter
// 订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetOrderFlag(_orderFlag string) error {
	r._orderFlag = _orderFlag
	r.Set("order_flag", _orderFlag)
	return nil
}

// Get OrderFlag Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetOrderFlag() string {
	return r._orderFlag
}

// Set is OrderCreateTime Setter
// 单据创建时间
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetOrderCreateTime(_orderCreateTime string) error {
	r._orderCreateTime = _orderCreateTime
	r.Set("order_create_time", _orderCreateTime)
	return nil
}

// Get OrderCreateTime Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetOrderCreateTime() string {
	return r._orderCreateTime
}

// Set is SupplierCode Setter
// 供应商编码，往来单位编码
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetSupplierCode(_supplierCode string) error {
	r._supplierCode = _supplierCode
	r.Set("supplier_code", _supplierCode)
	return nil
}

// Get SupplierCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetSupplierCode() string {
	return r._supplierCode
}

// Set is SupplierName Setter
// 供应商名称 ，往来单位名称
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetSupplierName(_supplierName string) error {
	r._supplierName = _supplierName
	r.Set("supplier_name", _supplierName)
	return nil
}

// Get SupplierName Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetSupplierName() string {
	return r._supplierName
}

// Set is TmsServiceCode Setter
// 配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetTmsServiceCode(_tmsServiceCode string) error {
	r._tmsServiceCode = _tmsServiceCode
	r.Set("tms_service_code", _tmsServiceCode)
	return nil
}

// Get TmsServiceCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetTmsServiceCode() string {
	return r._tmsServiceCode
}

// Set is TmsServiceName Setter
// 快递公司名称
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetTmsServiceName(_tmsServiceName string) error {
	r._tmsServiceName = _tmsServiceName
	r.Set("tms_service_name", _tmsServiceName)
	return nil
}

// Get TmsServiceName Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetTmsServiceName() string {
	return r._tmsServiceName
}

// Set is TmsOrderCode Setter
// 运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetTmsOrderCode(_tmsOrderCode string) error {
	r._tmsOrderCode = _tmsOrderCode
	r.Set("tms_order_code", _tmsOrderCode)
	return nil
}

// Get TmsOrderCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetTmsOrderCode() string {
	return r._tmsOrderCode
}

// Set is PrevOrderCode Setter
// 来源单据号，如销售退货时填充原销售订单号
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetPrevOrderCode(_prevOrderCode string) error {
	r._prevOrderCode = _prevOrderCode
	r.Set("prev_order_code", _prevOrderCode)
	return nil
}

// Get PrevOrderCode Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetPrevOrderCode() string {
	return r._prevOrderCode
}

// Set is ReturnReason Setter
// 销退时请提供退货的原因
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetReturnReason(_returnReason string) error {
	r._returnReason = _returnReason
	r.Set("return_reason", _returnReason)
	return nil
}

// Get ReturnReason Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetReturnReason() string {
	return r._returnReason
}

// Set is ExpectStartTime Setter
// 预期送达开始时间
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetExpectStartTime(_expectStartTime string) error {
	r._expectStartTime = _expectStartTime
	r.Set("expect_start_time", _expectStartTime)
	return nil
}

// Get ExpectStartTime Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetExpectStartTime() string {
	return r._expectStartTime
}

// Set is ExpectEndTime Setter
// 预期送达结束时间
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetExpectEndTime(_expectEndTime string) error {
	r._expectEndTime = _expectEndTime
	r.Set("expect_end_time", _expectEndTime)
	return nil
}

// Get ExpectEndTime Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetExpectEndTime() string {
	return r._expectEndTime
}

// Set is ReceiverInfo Setter
// 系统自动生成
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetReceiverInfo(_receiverInfo *Receiverinfowlbwmsstockinordernotifywl) error {
	r._receiverInfo = _receiverInfo
	r.Set("receiver_info", _receiverInfo)
	return nil
}

// Get ReceiverInfo Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetReceiverInfo() *Receiverinfowlbwmsstockinordernotifywl {
	return r._receiverInfo
}

// Set is SenderInfo Setter
// 系统自动生成
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetSenderInfo(_senderInfo *Senderinfowlbwmsstockinordernotifywl) error {
	r._senderInfo = _senderInfo
	r.Set("sender_info", _senderInfo)
	return nil
}

// Get SenderInfo Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetSenderInfo() *Senderinfowlbwmsstockinordernotifywl {
	return r._senderInfo
}

// Set is OrderItemList Setter
// 系统自动生成
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsstockinordernotifywl) error {
	r._orderItemList = _orderItemList
	r.Set("order_item_list", _orderItemList)
	return nil
}

// Get OrderItemList Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetOrderItemList() []Orderitemlistwlbwmsstockinordernotifywl {
	return r._orderItemList
}

// Set is ExtendFields Setter
// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// Get ExtendFields Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetExtendFields() string {
	return r._extendFields
}

// Set is Remark Setter
// 备注，销退入库订单需要留言备注填充到此字段
func (r *TaobaoWlbWmsStockInOrderNotifyAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoWlbWmsStockInOrderNotifyAPIRequest) GetRemark() string {
	return r._remark
}
