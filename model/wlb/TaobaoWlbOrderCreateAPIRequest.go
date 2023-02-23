package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderCreateAPIRequest 创建物流宝订单 API请求
// taobao.wlb.order.create
//
// 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
type TaobaoWlbOrderCreateAPIRequest struct {
	model.Params
	// 该字段暂时保留
	_attributes string
	// 订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。
	_orderItemList string
	// 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，<br/>: 是否可改配送方式  默认可更改<br/>11 CONSIGN 物流宝代理发货,自动更改发货状态<br/>12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
	_orderFlag string
	// 仓库编码
	_storeCode string
	// 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用
	_outBizCode string
	// 源订单编号
	_prevOrderCode string
	// 物流公司编码
	_tmsServiceCode string
	// 备注
	_remark string
	// 买家呢称
	_buyerNick string
	// 计划开始送达时间  在入库单中可能会使用
	_expectStartTime string
	// 期望结束时间，在入库单会使用到
	_expectEndTime string
	// 投递时间范围要求,格式'13:20'用分号隔开
	_scheduleStart string
	// 投递时间范围要求,格式'15:20'用分号隔开
	_scheduleEnd string
	// 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。
	_orderCode string
	// 运单编号，退货单时可能会使用
	_tmsOrderCode string
	// 订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库
	_orderType string
	// 订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购
	_orderSubType string
	// {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]}
	_invoinceInfo string
	// 出库单中可能会用到<br/>运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号<br/><br/>========================================<br/>如果某一个字段的数据为空时，必须传NA
	_tmsInfo string
	// 支付宝交易号
	_alipayNo string
	// 投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY
	_scheduleType string
	// 收货方信息，必须传， 手机和电话必选其一。收货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
	_receiverInfo string
	// 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
	_senderInfo string
	// 总金额
	_totalAmount int64
	// 应收金额，cod订单必选
	_payableAmount int64
	// cod服务费，只有cod订单的时候，才需要这个字段
	_serviceFee int64
	// 包裹件数，入库单和出库单中会用到
	_packageCount int64
	// 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。
	_isFinished bool
}

// NewTaobaoWlbOrderCreateRequest 初始化TaobaoWlbOrderCreateAPIRequest对象
func NewTaobaoWlbOrderCreateRequest() *TaobaoWlbOrderCreateAPIRequest {
	return &TaobaoWlbOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributes is Attributes Setter
// 该字段暂时保留
func (r *TaobaoWlbOrderCreateAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetOrderItemList is OrderItemList Setter
// 订单商品列表： {&#34;order_item_list&#34;:[{&#34;trade_code&#34;:&#34;可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号&#34;,&#34; sub_trade_code &#34;:&#34;可选,淘宝子交易号&#34;,&#34;item_id&#34;:&#34;必须,商品Id&#34;,&#34;item_code&#34;:&#34;必须,商家编码&#34;,&#34;item_name&#34;:&#34;可选,物流宝商品名称&#34;,&#34;item_quantity&#34;:&#34;必选,计划数量&#34;,&#34;item_price&#34;:&#34;必选,物品价格,单位为分&#34;,&#34;owner_user_nick&#34;:&#34;可选,货主nick 代销模式下会存在&#34;,&#34;flag&#34;:&#34;判断是否为赠品0 不是1是&#34;,&#34;remarks&#34;:&#34;可选,备注&#34;,&#34;batch_remark&#34;:&#34;可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联&#34;，&#34;inventory_type&#34;:&#34;库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存&#34;,&#34;picture_url&#34;:&#34;图片Url&#34;,&#34;distributor_user_nick&#34;: &#34;分销商NICK&#34;,必选&#34;ext_order_item_code&#34;:&#34;可选，外部商品的商家编码&#34;]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。
func (r *TaobaoWlbOrderCreateAPIRequest) SetOrderItemList(_orderItemList string) error {
	r._orderItemList = _orderItemList
	r.Set("order_item_list", _orderItemList)
	return nil
}

// GetOrderItemList OrderItemList Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetOrderItemList() string {
	return r._orderItemList
}

// SetOrderFlag is OrderFlag Setter
// 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，&lt;br/&gt;: 是否可改配送方式  默认可更改&lt;br/&gt;11 CONSIGN 物流宝代理发货,自动更改发货状态&lt;br/&gt;12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
func (r *TaobaoWlbOrderCreateAPIRequest) SetOrderFlag(_orderFlag string) error {
	r._orderFlag = _orderFlag
	r.Set("order_flag", _orderFlag)
	return nil
}

// GetOrderFlag OrderFlag Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetOrderFlag() string {
	return r._orderFlag
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbOrderCreateAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetOutBizCode is OutBizCode Setter
// 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用
func (r *TaobaoWlbOrderCreateAPIRequest) SetOutBizCode(_outBizCode string) error {
	r._outBizCode = _outBizCode
	r.Set("out_biz_code", _outBizCode)
	return nil
}

// GetOutBizCode OutBizCode Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetOutBizCode() string {
	return r._outBizCode
}

// SetPrevOrderCode is PrevOrderCode Setter
// 源订单编号
func (r *TaobaoWlbOrderCreateAPIRequest) SetPrevOrderCode(_prevOrderCode string) error {
	r._prevOrderCode = _prevOrderCode
	r.Set("prev_order_code", _prevOrderCode)
	return nil
}

// GetPrevOrderCode PrevOrderCode Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetPrevOrderCode() string {
	return r._prevOrderCode
}

// SetTmsServiceCode is TmsServiceCode Setter
// 物流公司编码
func (r *TaobaoWlbOrderCreateAPIRequest) SetTmsServiceCode(_tmsServiceCode string) error {
	r._tmsServiceCode = _tmsServiceCode
	r.Set("tms_service_code", _tmsServiceCode)
	return nil
}

// GetTmsServiceCode TmsServiceCode Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetTmsServiceCode() string {
	return r._tmsServiceCode
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoWlbOrderCreateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetRemark() string {
	return r._remark
}

// SetBuyerNick is BuyerNick Setter
// 买家呢称
func (r *TaobaoWlbOrderCreateAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetExpectStartTime is ExpectStartTime Setter
// 计划开始送达时间  在入库单中可能会使用
func (r *TaobaoWlbOrderCreateAPIRequest) SetExpectStartTime(_expectStartTime string) error {
	r._expectStartTime = _expectStartTime
	r.Set("expect_start_time", _expectStartTime)
	return nil
}

// GetExpectStartTime ExpectStartTime Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetExpectStartTime() string {
	return r._expectStartTime
}

// SetExpectEndTime is ExpectEndTime Setter
// 期望结束时间，在入库单会使用到
func (r *TaobaoWlbOrderCreateAPIRequest) SetExpectEndTime(_expectEndTime string) error {
	r._expectEndTime = _expectEndTime
	r.Set("expect_end_time", _expectEndTime)
	return nil
}

// GetExpectEndTime ExpectEndTime Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetExpectEndTime() string {
	return r._expectEndTime
}

// SetScheduleStart is ScheduleStart Setter
// 投递时间范围要求,格式&#39;13:20&#39;用分号隔开
func (r *TaobaoWlbOrderCreateAPIRequest) SetScheduleStart(_scheduleStart string) error {
	r._scheduleStart = _scheduleStart
	r.Set("schedule_start", _scheduleStart)
	return nil
}

// GetScheduleStart ScheduleStart Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetScheduleStart() string {
	return r._scheduleStart
}

// SetScheduleEnd is ScheduleEnd Setter
// 投递时间范围要求,格式&#39;15:20&#39;用分号隔开
func (r *TaobaoWlbOrderCreateAPIRequest) SetScheduleEnd(_scheduleEnd string) error {
	r._scheduleEnd = _scheduleEnd
	r.Set("schedule_end", _scheduleEnd)
	return nil
}

// GetScheduleEnd ScheduleEnd Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetScheduleEnd() string {
	return r._scheduleEnd
}

// SetOrderCode is OrderCode Setter
// 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。
func (r *TaobaoWlbOrderCreateAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetTmsOrderCode is TmsOrderCode Setter
// 运单编号，退货单时可能会使用
func (r *TaobaoWlbOrderCreateAPIRequest) SetTmsOrderCode(_tmsOrderCode string) error {
	r._tmsOrderCode = _tmsOrderCode
	r.Set("tms_order_code", _tmsOrderCode)
	return nil
}

// GetTmsOrderCode TmsOrderCode Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetTmsOrderCode() string {
	return r._tmsOrderCode
}

// SetOrderType is OrderType Setter
// 订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库
func (r *TaobaoWlbOrderCreateAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetOrderSubType is OrderSubType Setter
// 订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购
func (r *TaobaoWlbOrderCreateAPIRequest) SetOrderSubType(_orderSubType string) error {
	r._orderSubType = _orderSubType
	r.Set("order_sub_type", _orderSubType)
	return nil
}

// GetOrderSubType OrderSubType Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetOrderSubType() string {
	return r._orderSubType
}

// SetInvoinceInfo is InvoinceInfo Setter
// {&#34;invoince_info&#34;: [{&#34;bill_type&#34;:&#34;发票类型，必选&#34;, &#34;bill_title&#34;:&#34;发票抬头，必选&#34;, &#34;bill_amount&#34;:&#34;发票金额(单位是分)，必选&#34;,&#34;bill_content&#34;:&#34;发票内容，可选&#34;}]}
func (r *TaobaoWlbOrderCreateAPIRequest) SetInvoinceInfo(_invoinceInfo string) error {
	r._invoinceInfo = _invoinceInfo
	r.Set("invoince_info", _invoinceInfo)
	return nil
}

// GetInvoinceInfo InvoinceInfo Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetInvoinceInfo() string {
	return r._invoinceInfo
}

// SetTmsInfo is TmsInfo Setter
// 出库单中可能会用到&lt;br/&gt;运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号&lt;br/&gt;&lt;br/&gt;========================================&lt;br/&gt;如果某一个字段的数据为空时，必须传NA
func (r *TaobaoWlbOrderCreateAPIRequest) SetTmsInfo(_tmsInfo string) error {
	r._tmsInfo = _tmsInfo
	r.Set("tms_info", _tmsInfo)
	return nil
}

// GetTmsInfo TmsInfo Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetTmsInfo() string {
	return r._tmsInfo
}

// SetAlipayNo is AlipayNo Setter
// 支付宝交易号
func (r *TaobaoWlbOrderCreateAPIRequest) SetAlipayNo(_alipayNo string) error {
	r._alipayNo = _alipayNo
	r.Set("alipay_no", _alipayNo)
	return nil
}

// GetAlipayNo AlipayNo Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetAlipayNo() string {
	return r._alipayNo
}

// SetScheduleType is ScheduleType Setter
// 投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY
func (r *TaobaoWlbOrderCreateAPIRequest) SetScheduleType(_scheduleType string) error {
	r._scheduleType = _scheduleType
	r.Set("schedule_type", _scheduleType)
	return nil
}

// GetScheduleType ScheduleType Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetScheduleType() string {
	return r._scheduleType
}

// SetReceiverInfo is ReceiverInfo Setter
// 收货方信息，必须传， 手机和电话必选其一。收货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
func (r *TaobaoWlbOrderCreateAPIRequest) SetReceiverInfo(_receiverInfo string) error {
	r._receiverInfo = _receiverInfo
	r.Set("receiver_info", _receiverInfo)
	return nil
}

// GetReceiverInfo ReceiverInfo Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetReceiverInfo() string {
	return r._receiverInfo
}

// SetSenderInfo is SenderInfo Setter
// 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
func (r *TaobaoWlbOrderCreateAPIRequest) SetSenderInfo(_senderInfo string) error {
	r._senderInfo = _senderInfo
	r.Set("sender_info", _senderInfo)
	return nil
}

// GetSenderInfo SenderInfo Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetSenderInfo() string {
	return r._senderInfo
}

// SetTotalAmount is TotalAmount Setter
// 总金额
func (r *TaobaoWlbOrderCreateAPIRequest) SetTotalAmount(_totalAmount int64) error {
	r._totalAmount = _totalAmount
	r.Set("total_amount", _totalAmount)
	return nil
}

// GetTotalAmount TotalAmount Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetTotalAmount() int64 {
	return r._totalAmount
}

// SetPayableAmount is PayableAmount Setter
// 应收金额，cod订单必选
func (r *TaobaoWlbOrderCreateAPIRequest) SetPayableAmount(_payableAmount int64) error {
	r._payableAmount = _payableAmount
	r.Set("payable_amount", _payableAmount)
	return nil
}

// GetPayableAmount PayableAmount Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetPayableAmount() int64 {
	return r._payableAmount
}

// SetServiceFee is ServiceFee Setter
// cod服务费，只有cod订单的时候，才需要这个字段
func (r *TaobaoWlbOrderCreateAPIRequest) SetServiceFee(_serviceFee int64) error {
	r._serviceFee = _serviceFee
	r.Set("service_fee", _serviceFee)
	return nil
}

// GetServiceFee ServiceFee Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetServiceFee() int64 {
	return r._serviceFee
}

// SetPackageCount is PackageCount Setter
// 包裹件数，入库单和出库单中会用到
func (r *TaobaoWlbOrderCreateAPIRequest) SetPackageCount(_packageCount int64) error {
	r._packageCount = _packageCount
	r.Set("package_count", _packageCount)
	return nil
}

// GetPackageCount PackageCount Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetPackageCount() int64 {
	return r._packageCount
}

// SetIsFinished is IsFinished Setter
// 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。
func (r *TaobaoWlbOrderCreateAPIRequest) SetIsFinished(_isFinished bool) error {
	r._isFinished = _isFinished
	r.Set("is_finished", _isFinished)
	return nil
}

// GetIsFinished IsFinished Getter
func (r TaobaoWlbOrderCreateAPIRequest) GetIsFinished() bool {
	return r._isFinished
}
