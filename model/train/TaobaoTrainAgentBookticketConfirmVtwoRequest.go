package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功v2--增加鉴权校验 API请求
taobao.train.agent.bookticket.confirm.vtwo

火车票代理商接口——确认出票是否成功
*/
type TaobaoTrainAgentBookticketConfirmVtwoRequest struct {
    model.Params
    // 错误的子订单号123434,123432
    subOrderId   string
    // 是否支持在线退改签
    canChange   bool
    // 主订单id
    mainOrderId   int64
    // 是否成功
    status   bool
    // 代理商id
    agentId   int64
    // 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
    tickets   []string
    // 订单中包含的票数量
    ticketNum   int64
    // 出发时间
    depDate   string
    // 12306成功出票id
    ticket12306Id   string
    // 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
    failMsg   string
    // 支付宝交易流水号
    alipayTradeNo   string
    // 订单类型0 代购 1直购 3抢票
    orderType   int64
    // 支付宝账号
    alipayAccount   string
    // 到达时间
    arriveDate   string
    // 出发站
    fromStationName   string
    // 到达站
    toStationName   string
    // 检票口
    boardingGates   string
    // 送票上门预计派送时间
    expectDeliveryTime   string
    // 扩展字段
    extendParams   string
    // 票信息列表
    ticketInfoList   []OrderTicketInfo
}

// 初始化TaobaoTrainAgentBookticketConfirmVtwoRequest对象
func NewTaobaoTrainAgentBookticketConfirmVtwoRequest() *TaobaoTrainAgentBookticketConfirmVtwoRequest{
    return &TaobaoTrainAgentBookticketConfirmVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookticket.confirm.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubOrderId Setter
// 错误的子订单号123434,123432
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetSubOrderId(subOrderId string) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

// SubOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetSubOrderId() string {
    return r.subOrderId
}
// CanChange Setter
// 是否支持在线退改签
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetCanChange(canChange bool) error {
    r.canChange = canChange
    r.Set("can_change", canChange)
    return nil
}

// CanChange Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetCanChange() bool {
    return r.canChange
}
// MainOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// Status Setter
// 是否成功
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetStatus(status bool) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetStatus() bool {
    return r.status
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAgentId() int64 {
    return r.agentId
}
// Tickets Setter
// 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTickets(tickets []string) error {
    r.tickets = tickets
    r.Set("tickets", tickets)
    return nil
}

// Tickets Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTickets() []string {
    return r.tickets
}
// TicketNum Setter
// 订单中包含的票数量
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicketNum(ticketNum int64) error {
    r.ticketNum = ticketNum
    r.Set("ticket_num", ticketNum)
    return nil
}

// TicketNum Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicketNum() int64 {
    return r.ticketNum
}
// DepDate Setter
// 出发时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetDepDate(depDate string) error {
    r.depDate = depDate
    r.Set("dep_date", depDate)
    return nil
}

// DepDate Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetDepDate() string {
    return r.depDate
}
// Ticket12306Id Setter
// 12306成功出票id
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicket12306Id(ticket12306Id string) error {
    r.ticket12306Id = ticket12306Id
    r.Set("ticket_12306_id", ticket12306Id)
    return nil
}

// Ticket12306Id Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicket12306Id() string {
    return r.ticket12306Id
}
// FailMsg Setter
// 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetFailMsg(failMsg string) error {
    r.failMsg = failMsg
    r.Set("fail_msg", failMsg)
    return nil
}

// FailMsg Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetFailMsg() string {
    return r.failMsg
}
// AlipayTradeNo Setter
// 支付宝交易流水号
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAlipayTradeNo(alipayTradeNo string) error {
    r.alipayTradeNo = alipayTradeNo
    r.Set("alipay_trade_no", alipayTradeNo)
    return nil
}

// AlipayTradeNo Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAlipayTradeNo() string {
    return r.alipayTradeNo
}
// OrderType Setter
// 订单类型0 代购 1直购 3抢票
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetOrderType() int64 {
    return r.orderType
}
// AlipayAccount Setter
// 支付宝账号
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAlipayAccount(alipayAccount string) error {
    r.alipayAccount = alipayAccount
    r.Set("alipay_account", alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAlipayAccount() string {
    return r.alipayAccount
}
// ArriveDate Setter
// 到达时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetArriveDate(arriveDate string) error {
    r.arriveDate = arriveDate
    r.Set("arrive_date", arriveDate)
    return nil
}

// ArriveDate Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetArriveDate() string {
    return r.arriveDate
}
// FromStationName Setter
// 出发站
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetFromStationName(fromStationName string) error {
    r.fromStationName = fromStationName
    r.Set("from_station_name", fromStationName)
    return nil
}

// FromStationName Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetFromStationName() string {
    return r.fromStationName
}
// ToStationName Setter
// 到达站
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetToStationName(toStationName string) error {
    r.toStationName = toStationName
    r.Set("to_station_name", toStationName)
    return nil
}

// ToStationName Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetToStationName() string {
    return r.toStationName
}
// BoardingGates Setter
// 检票口
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetBoardingGates(boardingGates string) error {
    r.boardingGates = boardingGates
    r.Set("boarding_gates", boardingGates)
    return nil
}

// BoardingGates Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetBoardingGates() string {
    return r.boardingGates
}
// ExpectDeliveryTime Setter
// 送票上门预计派送时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetExpectDeliveryTime(expectDeliveryTime string) error {
    r.expectDeliveryTime = expectDeliveryTime
    r.Set("expect_delivery_time", expectDeliveryTime)
    return nil
}

// ExpectDeliveryTime Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetExpectDeliveryTime() string {
    return r.expectDeliveryTime
}
// ExtendParams Setter
// 扩展字段
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

// ExtendParams Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetExtendParams() string {
    return r.extendParams
}
// TicketInfoList Setter
// 票信息列表
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicketInfoList(ticketInfoList []OrderTicketInfo) error {
    r.ticketInfoList = ticketInfoList
    r.Set("ticket_info_list", ticketInfoList)
    return nil
}

// TicketInfoList Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicketInfoList() []OrderTicketInfo {
    return r.ticketInfoList
}
