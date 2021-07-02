package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsReturnOrderNotifyAPIRequest 销售退货通知 API请求
// taobao.wlb.wms.return.order.notify
//
// 销售退货通知
type TaobaoWlbWmsReturnOrderNotifyAPIRequest struct {
	model.Params
	// ERP单据编码
	_orderCode string
	// 仓库编码
	_storeCode string
	// 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
	_orderFlag string
	// 来源单据号，销售退货时填充原发货的LBX号
	_prevOrderCode string
	// 快递公司编码
	_tmsServiceCode string
	// 运单号
	_tmsOrderCode string
	// 销退时请提供退货的原因
	_returnReason string
	// 买家昵称
	_buyerNick string
	// 发件人信息
	_senderInfo *Senderinfowlbwmsreturnordernotify
	// 收件人信息
	_receiverInfo *Receiverinfowlbwmsreturnordernotify
	// 商品信息列表
	_orderItemList []Orderitemlistwlbwmsreturnordernotify
	// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
	_extendFields string
	// 备注
	_remark string
	// 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
	_orderSource string
	// 订单类型 501 销退入库
	_orderType string
	// 货主ID
	_ownerUserId string
	// ERP订单创建时间
	_orderCreateTime string
	// 快递公司名称
	_tmsServiceName string
}

// NewTaobaoWlbWmsReturnOrderNotifyRequest 初始化TaobaoWlbWmsReturnOrderNotifyAPIRequest对象
func NewTaobaoWlbWmsReturnOrderNotifyRequest() *TaobaoWlbWmsReturnOrderNotifyAPIRequest {
	return &TaobaoWlbWmsReturnOrderNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.return.order.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCode is OrderCode Setter
// ERP单据编码
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetOrderFlag is OrderFlag Setter
// 用字符串格式来表示订单标记列表：比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL 等，中间用“^”来隔开 ----------------------------------------  订单标记list（所有字母全部大写）： 9:VISIT-上门12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOrderFlag(_orderFlag string) error {
	r._orderFlag = _orderFlag
	r.Set("order_flag", _orderFlag)
	return nil
}

// GetOrderFlag OrderFlag Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOrderFlag() string {
	return r._orderFlag
}

// SetPrevOrderCode is PrevOrderCode Setter
// 来源单据号，销售退货时填充原发货的LBX号
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetPrevOrderCode(_prevOrderCode string) error {
	r._prevOrderCode = _prevOrderCode
	r.Set("prev_order_code", _prevOrderCode)
	return nil
}

// GetPrevOrderCode PrevOrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetPrevOrderCode() string {
	return r._prevOrderCode
}

// SetTmsServiceCode is TmsServiceCode Setter
// 快递公司编码
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetTmsServiceCode(_tmsServiceCode string) error {
	r._tmsServiceCode = _tmsServiceCode
	r.Set("tms_service_code", _tmsServiceCode)
	return nil
}

// GetTmsServiceCode TmsServiceCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetTmsServiceCode() string {
	return r._tmsServiceCode
}

// SetTmsOrderCode is TmsOrderCode Setter
// 运单号
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetTmsOrderCode(_tmsOrderCode string) error {
	r._tmsOrderCode = _tmsOrderCode
	r.Set("tms_order_code", _tmsOrderCode)
	return nil
}

// GetTmsOrderCode TmsOrderCode Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetTmsOrderCode() string {
	return r._tmsOrderCode
}

// SetReturnReason is ReturnReason Setter
// 销退时请提供退货的原因
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetReturnReason(_returnReason string) error {
	r._returnReason = _returnReason
	r.Set("return_reason", _returnReason)
	return nil
}

// GetReturnReason ReturnReason Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetReturnReason() string {
	return r._returnReason
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetSenderInfo is SenderInfo Setter
// 发件人信息
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetSenderInfo(_senderInfo *Senderinfowlbwmsreturnordernotify) error {
	r._senderInfo = _senderInfo
	r.Set("sender_info", _senderInfo)
	return nil
}

// GetSenderInfo SenderInfo Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetSenderInfo() *Senderinfowlbwmsreturnordernotify {
	return r._senderInfo
}

// SetReceiverInfo is ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetReceiverInfo(_receiverInfo *Receiverinfowlbwmsreturnordernotify) error {
	r._receiverInfo = _receiverInfo
	r.Set("receiver_info", _receiverInfo)
	return nil
}

// GetReceiverInfo ReceiverInfo Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetReceiverInfo() *Receiverinfowlbwmsreturnordernotify {
	return r._receiverInfo
}

// SetOrderItemList is OrderItemList Setter
// 商品信息列表
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsreturnordernotify) error {
	r._orderItemList = _orderItemList
	r.Set("order_item_list", _orderItemList)
	return nil
}

// GetOrderItemList OrderItemList Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOrderItemList() []Orderitemlistwlbwmsreturnordernotify {
	return r._orderItemList
}

// SetExtendFields is ExtendFields Setter
// 扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// GetExtendFields ExtendFields Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetExtendFields() string {
	return r._extendFields
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetRemark() string {
	return r._remark
}

// SetOrderSource is OrderSource Setter
// 订单来源 201 淘宝，202 1688，203 苏宁，204 亚马逊中国，205 当当，206 ebay，207 唯品会，208 1号店，209 国美，210 拍拍，211 聚美优品，212 乐峰，214 京东，301 其他
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOrderSource(_orderSource string) error {
	r._orderSource = _orderSource
	r.Set("order_source", _orderSource)
	return nil
}

// GetOrderSource OrderSource Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOrderSource() string {
	return r._orderSource
}

// SetOrderType is OrderType Setter
// 订单类型 501 销退入库
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetOwnerUserId is OwnerUserId Setter
// 货主ID
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOwnerUserId(_ownerUserId string) error {
	r._ownerUserId = _ownerUserId
	r.Set("owner_user_id", _ownerUserId)
	return nil
}

// GetOwnerUserId OwnerUserId Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOwnerUserId() string {
	return r._ownerUserId
}

// SetOrderCreateTime is OrderCreateTime Setter
// ERP订单创建时间
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetOrderCreateTime(_orderCreateTime string) error {
	r._orderCreateTime = _orderCreateTime
	r.Set("order_create_time", _orderCreateTime)
	return nil
}

// GetOrderCreateTime OrderCreateTime Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetOrderCreateTime() string {
	return r._orderCreateTime
}

// SetTmsServiceName is TmsServiceName Setter
// 快递公司名称
func (r *TaobaoWlbWmsReturnOrderNotifyAPIRequest) SetTmsServiceName(_tmsServiceName string) error {
	r._tmsServiceName = _tmsServiceName
	r.Set("tms_service_name", _tmsServiceName)
	return nil
}

// GetTmsServiceName TmsServiceName Getter
func (r TaobaoWlbWmsReturnOrderNotifyAPIRequest) GetTmsServiceName() string {
	return r._tmsServiceName
}
