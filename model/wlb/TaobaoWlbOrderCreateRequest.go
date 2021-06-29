package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建物流宝订单 API请求
taobao.wlb.order.create

创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
*/
type TaobaoWlbOrderCreateRequest struct {
    model.Params
    // 投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY
    _scheduleType   string
    // 订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购
    _orderSubType   string
    // 订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库
    _orderType   string
    // 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用
    _outBizCode   string
    // 仓库编码
    _storeCode   string
    // 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。
    _orderCode   string
    // 源订单编号
    _prevOrderCode   string
    // 物流公司编码
    _tmsServiceCode   string
    // 运单编号，退货单时可能会使用
    _tmsOrderCode   string
    // 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。
    _isFinished   bool
    // 投递时间范围要求,格式'13:20'用分号隔开
    _scheduleStart   string
    // 投递时间范围要求,格式'15:20'用分号隔开
    _scheduleEnd   string
    // 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，<br/>: 是否可改配送方式  默认可更改<br/>11 CONSIGN 物流宝代理发货,自动更改发货状态<br/>12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    _orderFlag   string
    // 支付宝交易号
    _alipayNo   string
    // 总金额
    _totalAmount   int64
    // 应收金额，cod订单必选
    _payableAmount   int64
    // cod服务费，只有cod订单的时候，才需要这个字段
    _serviceFee   int64
    // 买家呢称
    _buyerNick   string
    // 收货方信息，必须传， 手机和电话必选其一。收货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
    _receiverInfo   string
    // 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
    _senderInfo   string
    // 计划开始送达时间  在入库单中可能会使用
    _expectStartTime   string
    // 期望结束时间，在入库单会使用到
    _expectEndTime   string
    // 包裹件数，入库单和出库单中会用到
    _packageCount   int64
    // 出库单中可能会用到<br/>运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号<br/><br/>========================================<br/>如果某一个字段的数据为空时，必须传NA
    _tmsInfo   string
    // {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]}
    _invoinceInfo   string
    // 订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。
    _orderItemList   string
    // 该字段暂时保留
    _attributes   string
    // 备注
    _remark   string
}

// 初始化TaobaoWlbOrderCreateRequest对象
func NewTaobaoWlbOrderCreateRequest() *TaobaoWlbOrderCreateRequest{
    return &TaobaoWlbOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderCreateRequest) GetApiMethodName() string {
    return "taobao.wlb.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScheduleType Setter
// 投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY
func (r *TaobaoWlbOrderCreateRequest) SetScheduleType(_scheduleType string) error {
    r._scheduleType = _scheduleType
    r.Set("schedule_type", _scheduleType)
    return nil
}

// ScheduleType Getter
func (r TaobaoWlbOrderCreateRequest) GetScheduleType() string {
    return r._scheduleType
}
// OrderSubType Setter
// 订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购
func (r *TaobaoWlbOrderCreateRequest) SetOrderSubType(_orderSubType string) error {
    r._orderSubType = _orderSubType
    r.Set("order_sub_type", _orderSubType)
    return nil
}

// OrderSubType Getter
func (r TaobaoWlbOrderCreateRequest) GetOrderSubType() string {
    return r._orderSubType
}
// OrderType Setter
// 订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库
func (r *TaobaoWlbOrderCreateRequest) SetOrderType(_orderType string) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbOrderCreateRequest) GetOrderType() string {
    return r._orderType
}
// OutBizCode Setter
// 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用
func (r *TaobaoWlbOrderCreateRequest) SetOutBizCode(_outBizCode string) error {
    r._outBizCode = _outBizCode
    r.Set("out_biz_code", _outBizCode)
    return nil
}

// OutBizCode Getter
func (r TaobaoWlbOrderCreateRequest) GetOutBizCode() string {
    return r._outBizCode
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbOrderCreateRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbOrderCreateRequest) GetStoreCode() string {
    return r._storeCode
}
// OrderCode Setter
// 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。
func (r *TaobaoWlbOrderCreateRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbOrderCreateRequest) GetOrderCode() string {
    return r._orderCode
}
// PrevOrderCode Setter
// 源订单编号
func (r *TaobaoWlbOrderCreateRequest) SetPrevOrderCode(_prevOrderCode string) error {
    r._prevOrderCode = _prevOrderCode
    r.Set("prev_order_code", _prevOrderCode)
    return nil
}

// PrevOrderCode Getter
func (r TaobaoWlbOrderCreateRequest) GetPrevOrderCode() string {
    return r._prevOrderCode
}
// TmsServiceCode Setter
// 物流公司编码
func (r *TaobaoWlbOrderCreateRequest) SetTmsServiceCode(_tmsServiceCode string) error {
    r._tmsServiceCode = _tmsServiceCode
    r.Set("tms_service_code", _tmsServiceCode)
    return nil
}

// TmsServiceCode Getter
func (r TaobaoWlbOrderCreateRequest) GetTmsServiceCode() string {
    return r._tmsServiceCode
}
// TmsOrderCode Setter
// 运单编号，退货单时可能会使用
func (r *TaobaoWlbOrderCreateRequest) SetTmsOrderCode(_tmsOrderCode string) error {
    r._tmsOrderCode = _tmsOrderCode
    r.Set("tms_order_code", _tmsOrderCode)
    return nil
}

// TmsOrderCode Getter
func (r TaobaoWlbOrderCreateRequest) GetTmsOrderCode() string {
    return r._tmsOrderCode
}
// IsFinished Setter
// 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。
func (r *TaobaoWlbOrderCreateRequest) SetIsFinished(_isFinished bool) error {
    r._isFinished = _isFinished
    r.Set("is_finished", _isFinished)
    return nil
}

// IsFinished Getter
func (r TaobaoWlbOrderCreateRequest) GetIsFinished() bool {
    return r._isFinished
}
// ScheduleStart Setter
// 投递时间范围要求,格式'13:20'用分号隔开
func (r *TaobaoWlbOrderCreateRequest) SetScheduleStart(_scheduleStart string) error {
    r._scheduleStart = _scheduleStart
    r.Set("schedule_start", _scheduleStart)
    return nil
}

// ScheduleStart Getter
func (r TaobaoWlbOrderCreateRequest) GetScheduleStart() string {
    return r._scheduleStart
}
// ScheduleEnd Setter
// 投递时间范围要求,格式'15:20'用分号隔开
func (r *TaobaoWlbOrderCreateRequest) SetScheduleEnd(_scheduleEnd string) error {
    r._scheduleEnd = _scheduleEnd
    r.Set("schedule_end", _scheduleEnd)
    return nil
}

// ScheduleEnd Getter
func (r TaobaoWlbOrderCreateRequest) GetScheduleEnd() string {
    return r._scheduleEnd
}
// OrderFlag Setter
// 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，<br/>: 是否可改配送方式  默认可更改<br/>11 CONSIGN 物流宝代理发货,自动更改发货状态<br/>12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
func (r *TaobaoWlbOrderCreateRequest) SetOrderFlag(_orderFlag string) error {
    r._orderFlag = _orderFlag
    r.Set("order_flag", _orderFlag)
    return nil
}

// OrderFlag Getter
func (r TaobaoWlbOrderCreateRequest) GetOrderFlag() string {
    return r._orderFlag
}
// AlipayNo Setter
// 支付宝交易号
func (r *TaobaoWlbOrderCreateRequest) SetAlipayNo(_alipayNo string) error {
    r._alipayNo = _alipayNo
    r.Set("alipay_no", _alipayNo)
    return nil
}

// AlipayNo Getter
func (r TaobaoWlbOrderCreateRequest) GetAlipayNo() string {
    return r._alipayNo
}
// TotalAmount Setter
// 总金额
func (r *TaobaoWlbOrderCreateRequest) SetTotalAmount(_totalAmount int64) error {
    r._totalAmount = _totalAmount
    r.Set("total_amount", _totalAmount)
    return nil
}

// TotalAmount Getter
func (r TaobaoWlbOrderCreateRequest) GetTotalAmount() int64 {
    return r._totalAmount
}
// PayableAmount Setter
// 应收金额，cod订单必选
func (r *TaobaoWlbOrderCreateRequest) SetPayableAmount(_payableAmount int64) error {
    r._payableAmount = _payableAmount
    r.Set("payable_amount", _payableAmount)
    return nil
}

// PayableAmount Getter
func (r TaobaoWlbOrderCreateRequest) GetPayableAmount() int64 {
    return r._payableAmount
}
// ServiceFee Setter
// cod服务费，只有cod订单的时候，才需要这个字段
func (r *TaobaoWlbOrderCreateRequest) SetServiceFee(_serviceFee int64) error {
    r._serviceFee = _serviceFee
    r.Set("service_fee", _serviceFee)
    return nil
}

// ServiceFee Getter
func (r TaobaoWlbOrderCreateRequest) GetServiceFee() int64 {
    return r._serviceFee
}
// BuyerNick Setter
// 买家呢称
func (r *TaobaoWlbOrderCreateRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoWlbOrderCreateRequest) GetBuyerNick() string {
    return r._buyerNick
}
// ReceiverInfo Setter
// 收货方信息，必须传， 手机和电话必选其一。收货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
func (r *TaobaoWlbOrderCreateRequest) SetReceiverInfo(_receiverInfo string) error {
    r._receiverInfo = _receiverInfo
    r.Set("receiver_info", _receiverInfo)
    return nil
}

// ReceiverInfo Getter
func (r TaobaoWlbOrderCreateRequest) GetReceiverInfo() string {
    return r._receiverInfo
}
// SenderInfo Setter
// 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话如果某一个字段的数据为空时，必须传NA
func (r *TaobaoWlbOrderCreateRequest) SetSenderInfo(_senderInfo string) error {
    r._senderInfo = _senderInfo
    r.Set("sender_info", _senderInfo)
    return nil
}

// SenderInfo Getter
func (r TaobaoWlbOrderCreateRequest) GetSenderInfo() string {
    return r._senderInfo
}
// ExpectStartTime Setter
// 计划开始送达时间  在入库单中可能会使用
func (r *TaobaoWlbOrderCreateRequest) SetExpectStartTime(_expectStartTime string) error {
    r._expectStartTime = _expectStartTime
    r.Set("expect_start_time", _expectStartTime)
    return nil
}

// ExpectStartTime Getter
func (r TaobaoWlbOrderCreateRequest) GetExpectStartTime() string {
    return r._expectStartTime
}
// ExpectEndTime Setter
// 期望结束时间，在入库单会使用到
func (r *TaobaoWlbOrderCreateRequest) SetExpectEndTime(_expectEndTime string) error {
    r._expectEndTime = _expectEndTime
    r.Set("expect_end_time", _expectEndTime)
    return nil
}

// ExpectEndTime Getter
func (r TaobaoWlbOrderCreateRequest) GetExpectEndTime() string {
    return r._expectEndTime
}
// PackageCount Setter
// 包裹件数，入库单和出库单中会用到
func (r *TaobaoWlbOrderCreateRequest) SetPackageCount(_packageCount int64) error {
    r._packageCount = _packageCount
    r.Set("package_count", _packageCount)
    return nil
}

// PackageCount Getter
func (r TaobaoWlbOrderCreateRequest) GetPackageCount() int64 {
    return r._packageCount
}
// TmsInfo Setter
// 出库单中可能会用到<br/>运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号<br/><br/>========================================<br/>如果某一个字段的数据为空时，必须传NA
func (r *TaobaoWlbOrderCreateRequest) SetTmsInfo(_tmsInfo string) error {
    r._tmsInfo = _tmsInfo
    r.Set("tms_info", _tmsInfo)
    return nil
}

// TmsInfo Getter
func (r TaobaoWlbOrderCreateRequest) GetTmsInfo() string {
    return r._tmsInfo
}
// InvoinceInfo Setter
// {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]}
func (r *TaobaoWlbOrderCreateRequest) SetInvoinceInfo(_invoinceInfo string) error {
    r._invoinceInfo = _invoinceInfo
    r.Set("invoince_info", _invoinceInfo)
    return nil
}

// InvoinceInfo Getter
func (r TaobaoWlbOrderCreateRequest) GetInvoinceInfo() string {
    return r._invoinceInfo
}
// OrderItemList Setter
// 订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。
func (r *TaobaoWlbOrderCreateRequest) SetOrderItemList(_orderItemList string) error {
    r._orderItemList = _orderItemList
    r.Set("order_item_list", _orderItemList)
    return nil
}

// OrderItemList Getter
func (r TaobaoWlbOrderCreateRequest) GetOrderItemList() string {
    return r._orderItemList
}
// Attributes Setter
// 该字段暂时保留
func (r *TaobaoWlbOrderCreateRequest) SetAttributes(_attributes string) error {
    r._attributes = _attributes
    r.Set("attributes", _attributes)
    return nil
}

// Attributes Getter
func (r TaobaoWlbOrderCreateRequest) GetAttributes() string {
    return r._attributes
}
// Remark Setter
// 备注
func (r *TaobaoWlbOrderCreateRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoWlbOrderCreateRequest) GetRemark() string {
    return r._remark
}
