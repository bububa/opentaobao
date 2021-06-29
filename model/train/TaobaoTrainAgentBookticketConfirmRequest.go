package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功 API请求
taobao.train.agent.bookticket.confirm

火车票代理商接口——确认出票是否成功
*/
type TaobaoTrainAgentBookticketConfirmRequest struct {
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
    // 火车票子订单id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格 (座次号格式:坐席_车厢号_座位号，eg1:硬卧_09_03号下铺 eg2:硬座_02_03c ,注意:票价、保险价格必须到分)
    tickets   []string
    // 订单中包含的票数量
    ticketNum   int64
    // 出发时间
    depDate   string
    // 12306成功出票id
    ticket12306Id   string
    // 记录失败原因，传数字，1:票已售完,2:票价变动,3:乘车人已购相同车票,4:出票超时,5:乘车人证件未通过铁路局核验,6:发车时间变动,7:车次信息变更,8:12306故障,9:学生票信息有误,10:身份冒用,11:被限制高消费,12:坐票已售完,13:行程冲突,14:预售期变更,15:用户12306账号登录失败,16:12306账号存在未支付订单,17:用户常旅客已满,18:乘客信息有误,19:非法席别,20:车次停运,21:session登陆失败,22:账户已在其他地方登陆,23:帐号手机未核验,24:取消订单次数达到上限,25:帐号持有人身份未核验,26:邮寄地址无法保证及时送达,27:无法满足用户定制需求,28:您主动要求取消,0:未知原因
    failMsg   string
    // 支付宝交易流水号
    alipayTradeNo   string
    // 订单类型 0 代购 1直购 3抢票
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
}

// 初始化TaobaoTrainAgentBookticketConfirmRequest对象
func NewTaobaoTrainAgentBookticketConfirmRequest() *TaobaoTrainAgentBookticketConfirmRequest{
    return &TaobaoTrainAgentBookticketConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookticketConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookticket.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookticketConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubOrderId Setter
// 错误的子订单号123434,123432
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetSubOrderId(subOrderId string) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

// SubOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetSubOrderId() string {
    return r.subOrderId
}
// CanChange Setter
// 是否支持在线退改签
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetCanChange(canChange bool) error {
    r.canChange = canChange
    r.Set("can_change", canChange)
    return nil
}

// CanChange Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetCanChange() bool {
    return r.canChange
}
// MainOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// Status Setter
// 是否成功
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetStatus(status bool) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetStatus() bool {
    return r.status
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetAgentId() int64 {
    return r.agentId
}
// Tickets Setter
// 火车票子订单id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格 (座次号格式:坐席_车厢号_座位号，eg1:硬卧_09_03号下铺 eg2:硬座_02_03c ,注意:票价、保险价格必须到分)
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetTickets(tickets []string) error {
    r.tickets = tickets
    r.Set("tickets", tickets)
    return nil
}

// Tickets Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetTickets() []string {
    return r.tickets
}
// TicketNum Setter
// 订单中包含的票数量
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetTicketNum(ticketNum int64) error {
    r.ticketNum = ticketNum
    r.Set("ticket_num", ticketNum)
    return nil
}

// TicketNum Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetTicketNum() int64 {
    return r.ticketNum
}
// DepDate Setter
// 出发时间
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetDepDate(depDate string) error {
    r.depDate = depDate
    r.Set("dep_date", depDate)
    return nil
}

// DepDate Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetDepDate() string {
    return r.depDate
}
// Ticket12306Id Setter
// 12306成功出票id
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetTicket12306Id(ticket12306Id string) error {
    r.ticket12306Id = ticket12306Id
    r.Set("ticket_12306_id", ticket12306Id)
    return nil
}

// Ticket12306Id Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetTicket12306Id() string {
    return r.ticket12306Id
}
// FailMsg Setter
// 记录失败原因，传数字，1:票已售完,2:票价变动,3:乘车人已购相同车票,4:出票超时,5:乘车人证件未通过铁路局核验,6:发车时间变动,7:车次信息变更,8:12306故障,9:学生票信息有误,10:身份冒用,11:被限制高消费,12:坐票已售完,13:行程冲突,14:预售期变更,15:用户12306账号登录失败,16:12306账号存在未支付订单,17:用户常旅客已满,18:乘客信息有误,19:非法席别,20:车次停运,21:session登陆失败,22:账户已在其他地方登陆,23:帐号手机未核验,24:取消订单次数达到上限,25:帐号持有人身份未核验,26:邮寄地址无法保证及时送达,27:无法满足用户定制需求,28:您主动要求取消,0:未知原因
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetFailMsg(failMsg string) error {
    r.failMsg = failMsg
    r.Set("fail_msg", failMsg)
    return nil
}

// FailMsg Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetFailMsg() string {
    return r.failMsg
}
// AlipayTradeNo Setter
// 支付宝交易流水号
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetAlipayTradeNo(alipayTradeNo string) error {
    r.alipayTradeNo = alipayTradeNo
    r.Set("alipay_trade_no", alipayTradeNo)
    return nil
}

// AlipayTradeNo Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetAlipayTradeNo() string {
    return r.alipayTradeNo
}
// OrderType Setter
// 订单类型 0 代购 1直购 3抢票
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetOrderType() int64 {
    return r.orderType
}
// AlipayAccount Setter
// 支付宝账号
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetAlipayAccount(alipayAccount string) error {
    r.alipayAccount = alipayAccount
    r.Set("alipay_account", alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetAlipayAccount() string {
    return r.alipayAccount
}
// ArriveDate Setter
// 到达时间
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetArriveDate(arriveDate string) error {
    r.arriveDate = arriveDate
    r.Set("arrive_date", arriveDate)
    return nil
}

// ArriveDate Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetArriveDate() string {
    return r.arriveDate
}
// FromStationName Setter
// 出发站
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetFromStationName(fromStationName string) error {
    r.fromStationName = fromStationName
    r.Set("from_station_name", fromStationName)
    return nil
}

// FromStationName Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetFromStationName() string {
    return r.fromStationName
}
// ToStationName Setter
// 到达站
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetToStationName(toStationName string) error {
    r.toStationName = toStationName
    r.Set("to_station_name", toStationName)
    return nil
}

// ToStationName Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetToStationName() string {
    return r.toStationName
}
// BoardingGates Setter
// 检票口
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetBoardingGates(boardingGates string) error {
    r.boardingGates = boardingGates
    r.Set("boarding_gates", boardingGates)
    return nil
}

// BoardingGates Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetBoardingGates() string {
    return r.boardingGates
}
