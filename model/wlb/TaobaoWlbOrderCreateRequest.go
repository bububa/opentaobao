package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建物流宝订单 APIRequest
taobao.wlb.order.create

创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
*/
type TaobaoWlbOrderCreateRequest struct {
    model.Params

    // 投递时延要求:  （1）INSTANT_ARRIVED： 当日达  （2）TOMMORROY_MORNING_ARRIVED：次晨达  （3）TOMMORROY_ARRIVED：次日达  （4）工作日：WORK_DAY  （5）节假日：WEEKED_DAY
    scheduleType   string 

    // 订单子类型：  （1）OTHER： 其他  （2）TAOBAO_TRADE： 淘宝交易  （3）OTHER_TRADE：其他交易  （4）ALLCOATE： 调拨  （5）PURCHASE:采购
    orderSubType   string 

    // 订单类型:  （1）NORMAL_OUT ：正常出库  （2）NORMAL_IN：正常入库  （3）RETURN_IN：退货入库  （4）EXCHANGE_OUT：换货出库
    orderType   string 

    // 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用
    outBizCode   string 

    // 仓库编码
    storeCode   string 

    // 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。
    orderCode   string 

    // 源订单编号
    prevOrderCode   string 

    // 物流公司编码
    tmsServiceCode   string 

    // 运单编号，退货单时可能会使用
    tmsOrderCode   string 

    // 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。
    isFinished   bool 

    // 投递时间范围要求,格式'13:20'用分号隔开
    scheduleStart   string 

    // 投递时间范围要求,格式'15:20'用分号隔开
    scheduleEnd   string 

    // 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，<br/>: 是否可改配送方式  默认可更改<br/>11 CONSIGN 物流宝代理发货,自动更改发货状态<br/>12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票
    orderFlag   string 

    // 支付宝交易号
    alipayNo   string 

    // 总金额
    totalAmount   int64 

    // 应收金额，cod订单必选
    payableAmount   int64 

    // cod服务费，只有cod订单的时候，才需要这个字段
    serviceFee   int64 

    // 买家呢称
    buyerNick   string 

    // 收货方信息，必须传， 手机和电话必选其一。
收货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话

如果某一个字段的数据为空时，必须传NA
    receiverInfo   string 

    // 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
如果某一个字段的数据为空时，必须传NA
    senderInfo   string 

    // 计划开始送达时间  在入库单中可能会使用
    expectStartTime   string 

    // 期望结束时间，在入库单会使用到
    expectEndTime   string 

    // 包裹件数，入库单和出库单中会用到
    packageCount   int64 

    // 出库单中可能会用到<br/>运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号<br/><br/>========================================<br/>如果某一个字段的数据为空时，必须传NA
    tmsInfo   string 

    // {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]}
    invoinceInfo   string 

    // 订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick
":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。
    orderItemList   string 

    // 该字段暂时保留
    attributes   string 

    // 备注
    remark   string 

    // 备注
    remark   string 

}

func NewTaobaoWlbOrderCreateRequest() *TaobaoWlbOrderCreateRequest{
    return &TaobaoWlbOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderCreateRequest) GetApiMethodName() string {
    return "taobao.wlb.order.create"
}

func (r TaobaoWlbOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderCreateRequest) SetScheduleType(scheduleType string) error {
    r.scheduleType = scheduleType
    r.Set("schedule_type", scheduleType)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetScheduleType() string {
    return r.scheduleType
}

func (r *TaobaoWlbOrderCreateRequest) SetOrderSubType(orderSubType string) error {
    r.orderSubType = orderSubType
    r.Set("order_sub_type", orderSubType)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetOrderSubType() string {
    return r.orderSubType
}

func (r *TaobaoWlbOrderCreateRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetOrderType() string {
    return r.orderType
}

func (r *TaobaoWlbOrderCreateRequest) SetOutBizCode(outBizCode string) error {
    r.outBizCode = outBizCode
    r.Set("out_biz_code", outBizCode)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetOutBizCode() string {
    return r.outBizCode
}

func (r *TaobaoWlbOrderCreateRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbOrderCreateRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbOrderCreateRequest) SetPrevOrderCode(prevOrderCode string) error {
    r.prevOrderCode = prevOrderCode
    r.Set("prev_order_code", prevOrderCode)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetPrevOrderCode() string {
    return r.prevOrderCode
}

func (r *TaobaoWlbOrderCreateRequest) SetTmsServiceCode(tmsServiceCode string) error {
    r.tmsServiceCode = tmsServiceCode
    r.Set("tms_service_code", tmsServiceCode)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetTmsServiceCode() string {
    return r.tmsServiceCode
}

func (r *TaobaoWlbOrderCreateRequest) SetTmsOrderCode(tmsOrderCode string) error {
    r.tmsOrderCode = tmsOrderCode
    r.Set("tms_order_code", tmsOrderCode)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetTmsOrderCode() string {
    return r.tmsOrderCode
}

func (r *TaobaoWlbOrderCreateRequest) SetIsFinished(isFinished bool) error {
    r.isFinished = isFinished
    r.Set("is_finished", isFinished)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetIsFinished() bool {
    return r.isFinished
}

func (r *TaobaoWlbOrderCreateRequest) SetScheduleStart(scheduleStart string) error {
    r.scheduleStart = scheduleStart
    r.Set("schedule_start", scheduleStart)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetScheduleStart() string {
    return r.scheduleStart
}

func (r *TaobaoWlbOrderCreateRequest) SetScheduleEnd(scheduleEnd string) error {
    r.scheduleEnd = scheduleEnd
    r.Set("schedule_end", scheduleEnd)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetScheduleEnd() string {
    return r.scheduleEnd
}

func (r *TaobaoWlbOrderCreateRequest) SetOrderFlag(orderFlag string) error {
    r.orderFlag = orderFlag
    r.Set("order_flag", orderFlag)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetOrderFlag() string {
    return r.orderFlag
}

func (r *TaobaoWlbOrderCreateRequest) SetAlipayNo(alipayNo string) error {
    r.alipayNo = alipayNo
    r.Set("alipay_no", alipayNo)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetAlipayNo() string {
    return r.alipayNo
}

func (r *TaobaoWlbOrderCreateRequest) SetTotalAmount(totalAmount int64) error {
    r.totalAmount = totalAmount
    r.Set("total_amount", totalAmount)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetTotalAmount() int64 {
    return r.totalAmount
}

func (r *TaobaoWlbOrderCreateRequest) SetPayableAmount(payableAmount int64) error {
    r.payableAmount = payableAmount
    r.Set("payable_amount", payableAmount)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetPayableAmount() int64 {
    return r.payableAmount
}

func (r *TaobaoWlbOrderCreateRequest) SetServiceFee(serviceFee int64) error {
    r.serviceFee = serviceFee
    r.Set("service_fee", serviceFee)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetServiceFee() int64 {
    return r.serviceFee
}

func (r *TaobaoWlbOrderCreateRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TaobaoWlbOrderCreateRequest) SetReceiverInfo(receiverInfo string) error {
    r.receiverInfo = receiverInfo
    r.Set("receiver_info", receiverInfo)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetReceiverInfo() string {
    return r.receiverInfo
}

func (r *TaobaoWlbOrderCreateRequest) SetSenderInfo(senderInfo string) error {
    r.senderInfo = senderInfo
    r.Set("sender_info", senderInfo)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetSenderInfo() string {
    return r.senderInfo
}

func (r *TaobaoWlbOrderCreateRequest) SetExpectStartTime(expectStartTime string) error {
    r.expectStartTime = expectStartTime
    r.Set("expect_start_time", expectStartTime)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetExpectStartTime() string {
    return r.expectStartTime
}

func (r *TaobaoWlbOrderCreateRequest) SetExpectEndTime(expectEndTime string) error {
    r.expectEndTime = expectEndTime
    r.Set("expect_end_time", expectEndTime)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetExpectEndTime() string {
    return r.expectEndTime
}

func (r *TaobaoWlbOrderCreateRequest) SetPackageCount(packageCount int64) error {
    r.packageCount = packageCount
    r.Set("package_count", packageCount)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetPackageCount() int64 {
    return r.packageCount
}

func (r *TaobaoWlbOrderCreateRequest) SetTmsInfo(tmsInfo string) error {
    r.tmsInfo = tmsInfo
    r.Set("tms_info", tmsInfo)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetTmsInfo() string {
    return r.tmsInfo
}

func (r *TaobaoWlbOrderCreateRequest) SetInvoinceInfo(invoinceInfo string) error {
    r.invoinceInfo = invoinceInfo
    r.Set("invoince_info", invoinceInfo)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetInvoinceInfo() string {
    return r.invoinceInfo
}

func (r *TaobaoWlbOrderCreateRequest) SetOrderItemList(orderItemList string) error {
    r.orderItemList = orderItemList
    r.Set("order_item_list", orderItemList)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetOrderItemList() string {
    return r.orderItemList
}

func (r *TaobaoWlbOrderCreateRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetAttributes() string {
    return r.attributes
}

func (r *TaobaoWlbOrderCreateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoWlbOrderCreateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoWlbOrderCreateRequest) GetRemark() string {
    return r.remark
}

