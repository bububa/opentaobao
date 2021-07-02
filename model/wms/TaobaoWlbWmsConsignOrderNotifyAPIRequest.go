package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsConsignOrderNotifyAPIRequest 发货订单通知 API请求
// taobao.wlb.wms.consign.order.notify
//
// 发货订单通知
type TaobaoWlbWmsConsignOrderNotifyAPIRequest struct {
	model.Params
	// ERP订单号
	_orderCode string
	// 单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单
	_orderType int64
	// 订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票)
	_orderFlag string
	// 订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
	_orderSource int64
	// 仓库编码，此字段为空时，由菜鸟路由仓库发货
	_storeCode string
	// 快递公司编码，此字段为空时，由菜鸟选择快递配送
	_tmsServiceCode string
	// 快递公司名称
	_tmsServiceName string
	// 前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容
	_prevOrderCode string
	// 下单时间，订单在交易平台创建时间
	_orderShopCreateTime string
	// 订单支付时间
	_orderPayTime string
	// 订单创建时间,ERP创建订单时间
	_orderCreateTime string
	// 订单审核时间,ERP创建支付时间
	_orderExaminationTime string
	// 订单总金额,=总商品金额-订单优惠金额+快递费用，单位分
	_orderAmount int64
	// 订单优惠金额，整单优惠金额，单位分
	_discountAmount int64
	// 订单应收金额，消费者还需要付的金额，单位分
	_arAmount int64
	// 订单已付金额，消费者已经支付的金额，单位分
	_gotAmount int64
	// 快递费用，单位分
	_postfee int64
	// COD服务费，单位分
	_serviceFee int64
	// 配送要求
	_deliverRequirements *Deliverrequirementswlbwmsconsignordernotify
	// 收件人信息
	_receiverInfo *Receiverwlbwmsconsignordernotify
	// 发货方信息
	_senderInfo *Senderwlbwmsconsignordernotify
	// 订单商品信息列表
	_orderItemList []Orderitemlistwlbwmsconsignordernotify
	// 发票信息列表
	_invoiceInfoList []Invoicelistwlbwmsconsignordernotify
	// 拓展属性
	_extendFields string
	// 备注
	_remark string
}

// NewTaobaoWlbWmsConsignOrderNotifyRequest 初始化TaobaoWlbWmsConsignOrderNotifyAPIRequest对象
func NewTaobaoWlbWmsConsignOrderNotifyRequest() *TaobaoWlbWmsConsignOrderNotifyAPIRequest {
	return &TaobaoWlbWmsConsignOrderNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.consign.order.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCode is OrderCode Setter
// ERP订单号
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetOrderType is OrderType Setter
// 单据类型 201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// SetOrderFlag is OrderFlag Setter
// 订单标识 (1: cod –货到付款，4:invoiceinfo-需要发票)
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderFlag(_orderFlag string) error {
	r._orderFlag = _orderFlag
	r.Set("order_flag", _orderFlag)
	return nil
}

// GetOrderFlag OrderFlag Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderFlag() string {
	return r._orderFlag
}

// SetOrderSource is OrderSource Setter
// 订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderSource(_orderSource int64) error {
	r._orderSource = _orderSource
	r.Set("order_source", _orderSource)
	return nil
}

// GetOrderSource OrderSource Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderSource() int64 {
	return r._orderSource
}

// SetStoreCode is StoreCode Setter
// 仓库编码，此字段为空时，由菜鸟路由仓库发货
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetTmsServiceCode is TmsServiceCode Setter
// 快递公司编码，此字段为空时，由菜鸟选择快递配送
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetTmsServiceCode(_tmsServiceCode string) error {
	r._tmsServiceCode = _tmsServiceCode
	r.Set("tms_service_code", _tmsServiceCode)
	return nil
}

// GetTmsServiceCode TmsServiceCode Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetTmsServiceCode() string {
	return r._tmsServiceCode
}

// SetTmsServiceName is TmsServiceName Setter
// 快递公司名称
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetTmsServiceName(_tmsServiceName string) error {
	r._tmsServiceName = _tmsServiceName
	r.Set("tms_service_name", _tmsServiceName)
	return nil
}

// GetTmsServiceName TmsServiceName Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetTmsServiceName() string {
	return r._tmsServiceName
}

// SetPrevOrderCode is PrevOrderCode Setter
// 前物流订单号，订单类型为502 换货出库单 503 补发出库单时，需求传入此内容
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetPrevOrderCode(_prevOrderCode string) error {
	r._prevOrderCode = _prevOrderCode
	r.Set("prev_order_code", _prevOrderCode)
	return nil
}

// GetPrevOrderCode PrevOrderCode Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetPrevOrderCode() string {
	return r._prevOrderCode
}

// SetOrderShopCreateTime is OrderShopCreateTime Setter
// 下单时间，订单在交易平台创建时间
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderShopCreateTime(_orderShopCreateTime string) error {
	r._orderShopCreateTime = _orderShopCreateTime
	r.Set("order_shop_create_time", _orderShopCreateTime)
	return nil
}

// GetOrderShopCreateTime OrderShopCreateTime Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderShopCreateTime() string {
	return r._orderShopCreateTime
}

// SetOrderPayTime is OrderPayTime Setter
// 订单支付时间
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderPayTime(_orderPayTime string) error {
	r._orderPayTime = _orderPayTime
	r.Set("order_pay_time", _orderPayTime)
	return nil
}

// GetOrderPayTime OrderPayTime Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderPayTime() string {
	return r._orderPayTime
}

// SetOrderCreateTime is OrderCreateTime Setter
// 订单创建时间,ERP创建订单时间
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderCreateTime(_orderCreateTime string) error {
	r._orderCreateTime = _orderCreateTime
	r.Set("order_create_time", _orderCreateTime)
	return nil
}

// GetOrderCreateTime OrderCreateTime Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderCreateTime() string {
	return r._orderCreateTime
}

// SetOrderExaminationTime is OrderExaminationTime Setter
// 订单审核时间,ERP创建支付时间
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderExaminationTime(_orderExaminationTime string) error {
	r._orderExaminationTime = _orderExaminationTime
	r.Set("order_examination_time", _orderExaminationTime)
	return nil
}

// GetOrderExaminationTime OrderExaminationTime Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderExaminationTime() string {
	return r._orderExaminationTime
}

// SetOrderAmount is OrderAmount Setter
// 订单总金额,=总商品金额-订单优惠金额+快递费用，单位分
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderAmount(_orderAmount int64) error {
	r._orderAmount = _orderAmount
	r.Set("order_amount", _orderAmount)
	return nil
}

// GetOrderAmount OrderAmount Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderAmount() int64 {
	return r._orderAmount
}

// SetDiscountAmount is DiscountAmount Setter
// 订单优惠金额，整单优惠金额，单位分
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetDiscountAmount(_discountAmount int64) error {
	r._discountAmount = _discountAmount
	r.Set("discount_amount", _discountAmount)
	return nil
}

// GetDiscountAmount DiscountAmount Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetDiscountAmount() int64 {
	return r._discountAmount
}

// SetArAmount is ArAmount Setter
// 订单应收金额，消费者还需要付的金额，单位分
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetArAmount(_arAmount int64) error {
	r._arAmount = _arAmount
	r.Set("ar_amount", _arAmount)
	return nil
}

// GetArAmount ArAmount Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetArAmount() int64 {
	return r._arAmount
}

// SetGotAmount is GotAmount Setter
// 订单已付金额，消费者已经支付的金额，单位分
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetGotAmount(_gotAmount int64) error {
	r._gotAmount = _gotAmount
	r.Set("got_amount", _gotAmount)
	return nil
}

// GetGotAmount GotAmount Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetGotAmount() int64 {
	return r._gotAmount
}

// SetPostfee is Postfee Setter
// 快递费用，单位分
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetPostfee(_postfee int64) error {
	r._postfee = _postfee
	r.Set("postfee", _postfee)
	return nil
}

// GetPostfee Postfee Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetPostfee() int64 {
	return r._postfee
}

// SetServiceFee is ServiceFee Setter
// COD服务费，单位分
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetServiceFee(_serviceFee int64) error {
	r._serviceFee = _serviceFee
	r.Set("service_fee", _serviceFee)
	return nil
}

// GetServiceFee ServiceFee Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetServiceFee() int64 {
	return r._serviceFee
}

// SetDeliverRequirements is DeliverRequirements Setter
// 配送要求
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetDeliverRequirements(_deliverRequirements *Deliverrequirementswlbwmsconsignordernotify) error {
	r._deliverRequirements = _deliverRequirements
	r.Set("deliver_requirements", _deliverRequirements)
	return nil
}

// GetDeliverRequirements DeliverRequirements Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetDeliverRequirements() *Deliverrequirementswlbwmsconsignordernotify {
	return r._deliverRequirements
}

// SetReceiverInfo is ReceiverInfo Setter
// 收件人信息
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetReceiverInfo(_receiverInfo *Receiverwlbwmsconsignordernotify) error {
	r._receiverInfo = _receiverInfo
	r.Set("receiver_info", _receiverInfo)
	return nil
}

// GetReceiverInfo ReceiverInfo Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetReceiverInfo() *Receiverwlbwmsconsignordernotify {
	return r._receiverInfo
}

// SetSenderInfo is SenderInfo Setter
// 发货方信息
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetSenderInfo(_senderInfo *Senderwlbwmsconsignordernotify) error {
	r._senderInfo = _senderInfo
	r.Set("sender_info", _senderInfo)
	return nil
}

// GetSenderInfo SenderInfo Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetSenderInfo() *Senderwlbwmsconsignordernotify {
	return r._senderInfo
}

// SetOrderItemList is OrderItemList Setter
// 订单商品信息列表
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetOrderItemList(_orderItemList []Orderitemlistwlbwmsconsignordernotify) error {
	r._orderItemList = _orderItemList
	r.Set("order_item_list", _orderItemList)
	return nil
}

// GetOrderItemList OrderItemList Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetOrderItemList() []Orderitemlistwlbwmsconsignordernotify {
	return r._orderItemList
}

// SetInvoiceInfoList is InvoiceInfoList Setter
// 发票信息列表
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetInvoiceInfoList(_invoiceInfoList []Invoicelistwlbwmsconsignordernotify) error {
	r._invoiceInfoList = _invoiceInfoList
	r.Set("invoice_info_list", _invoiceInfoList)
	return nil
}

// GetInvoiceInfoList InvoiceInfoList Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetInvoiceInfoList() []Invoicelistwlbwmsconsignordernotify {
	return r._invoiceInfoList
}

// SetExtendFields is ExtendFields Setter
// 拓展属性
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetExtendFields(_extendFields string) error {
	r._extendFields = _extendFields
	r.Set("extend_fields", _extendFields)
	return nil
}

// GetExtendFields ExtendFields Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetExtendFields() string {
	return r._extendFields
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoWlbWmsConsignOrderNotifyAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoWlbWmsConsignOrderNotifyAPIRequest) GetRemark() string {
	return r._remark
}
