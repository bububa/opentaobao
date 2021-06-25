package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认出票是否成功v2--增加鉴权校验 APIRequest
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
    tickets   []String 

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

func NewTaobaoTrainAgentBookticketConfirmVtwoRequest() *TaobaoTrainAgentBookticketConfirmVtwoRequest{
    return &TaobaoTrainAgentBookticketConfirmVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.bookticket.confirm.vtwo"
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetSubOrderId(subOrderId string) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetSubOrderId() string {
    return r.subOrderId
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetCanChange(canChange bool) error {
    r.canChange = canChange
    r.Set("can_change", canChange)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetCanChange() bool {
    return r.canChange
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetStatus(status bool) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetStatus() bool {
    return r.status
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTickets(tickets []String) error {
    r.tickets = tickets
    r.Set("tickets", tickets)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTickets() []String {
    return r.tickets
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicketNum(ticketNum int64) error {
    r.ticketNum = ticketNum
    r.Set("ticket_num", ticketNum)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicketNum() int64 {
    return r.ticketNum
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetDepDate(depDate string) error {
    r.depDate = depDate
    r.Set("dep_date", depDate)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetDepDate() string {
    return r.depDate
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicket12306Id(ticket12306Id string) error {
    r.ticket12306Id = ticket12306Id
    r.Set("ticket_12306_id", ticket12306Id)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicket12306Id() string {
    return r.ticket12306Id
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetFailMsg(failMsg string) error {
    r.failMsg = failMsg
    r.Set("fail_msg", failMsg)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetFailMsg() string {
    return r.failMsg
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAlipayTradeNo(alipayTradeNo string) error {
    r.alipayTradeNo = alipayTradeNo
    r.Set("alipay_trade_no", alipayTradeNo)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAlipayTradeNo() string {
    return r.alipayTradeNo
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAlipayAccount(alipayAccount string) error {
    r.alipayAccount = alipayAccount
    r.Set("alipay_account", alipayAccount)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAlipayAccount() string {
    return r.alipayAccount
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetArriveDate(arriveDate string) error {
    r.arriveDate = arriveDate
    r.Set("arrive_date", arriveDate)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetArriveDate() string {
    return r.arriveDate
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetFromStationName(fromStationName string) error {
    r.fromStationName = fromStationName
    r.Set("from_station_name", fromStationName)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetFromStationName() string {
    return r.fromStationName
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetToStationName(toStationName string) error {
    r.toStationName = toStationName
    r.Set("to_station_name", toStationName)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetToStationName() string {
    return r.toStationName
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetBoardingGates(boardingGates string) error {
    r.boardingGates = boardingGates
    r.Set("boarding_gates", boardingGates)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetBoardingGates() string {
    return r.boardingGates
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetExpectDeliveryTime(expectDeliveryTime string) error {
    r.expectDeliveryTime = expectDeliveryTime
    r.Set("expect_delivery_time", expectDeliveryTime)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetExpectDeliveryTime() string {
    return r.expectDeliveryTime
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetExtendParams() string {
    return r.extendParams
}

func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicketInfoList(ticketInfoList []OrderTicketInfo) error {
    r.ticketInfoList = ticketInfoList
    r.Set("ticket_info_list", ticketInfoList)
    return nil
}

func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicketInfoList() []OrderTicketInfo {
    return r.ticketInfoList
}

