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
    _subOrderId   string
    // 是否支持在线退改签
    _canChange   bool
    // 主订单id
    _mainOrderId   int64
    // 是否成功
    _status   bool
    // 代理商id
    _agentId   int64
    // 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
    _tickets   []string
    // 订单中包含的票数量
    _ticketNum   int64
    // 出发时间
    _depDate   string
    // 12306成功出票id
    _ticket12306Id   string
    // 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
    _failMsg   string
    // 支付宝交易流水号
    _alipayTradeNo   string
    // 订单类型0 代购 1直购 3抢票
    _orderType   int64
    // 支付宝账号
    _alipayAccount   string
    // 到达时间
    _arriveDate   string
    // 出发站
    _fromStationName   string
    // 到达站
    _toStationName   string
    // 检票口
    _boardingGates   string
    // 送票上门预计派送时间
    _expectDeliveryTime   string
    // 扩展字段
    _extendParams   string
    // 票信息列表
    _ticketInfoList   []OrderTicketInfo
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
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetSubOrderId(_subOrderId string) error {
    r._subOrderId = _subOrderId
    r.Set("sub_order_id", _subOrderId)
    return nil
}

// SubOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetSubOrderId() string {
    return r._subOrderId
}
// CanChange Setter
// 是否支持在线退改签
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetCanChange(_canChange bool) error {
    r._canChange = _canChange
    r.Set("can_change", _canChange)
    return nil
}

// CanChange Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetCanChange() bool {
    return r._canChange
}
// MainOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// Status Setter
// 是否成功
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetStatus(_status bool) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetStatus() bool {
    return r._status
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAgentId() int64 {
    return r._agentId
}
// Tickets Setter
// 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTickets(_tickets []string) error {
    r._tickets = _tickets
    r.Set("tickets", _tickets)
    return nil
}

// Tickets Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTickets() []string {
    return r._tickets
}
// TicketNum Setter
// 订单中包含的票数量
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicketNum(_ticketNum int64) error {
    r._ticketNum = _ticketNum
    r.Set("ticket_num", _ticketNum)
    return nil
}

// TicketNum Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicketNum() int64 {
    return r._ticketNum
}
// DepDate Setter
// 出发时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetDepDate(_depDate string) error {
    r._depDate = _depDate
    r.Set("dep_date", _depDate)
    return nil
}

// DepDate Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetDepDate() string {
    return r._depDate
}
// Ticket12306Id Setter
// 12306成功出票id
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicket12306Id(_ticket12306Id string) error {
    r._ticket12306Id = _ticket12306Id
    r.Set("ticket_12306_id", _ticket12306Id)
    return nil
}

// Ticket12306Id Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicket12306Id() string {
    return r._ticket12306Id
}
// FailMsg Setter
// 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetFailMsg(_failMsg string) error {
    r._failMsg = _failMsg
    r.Set("fail_msg", _failMsg)
    return nil
}

// FailMsg Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetFailMsg() string {
    return r._failMsg
}
// AlipayTradeNo Setter
// 支付宝交易流水号
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAlipayTradeNo(_alipayTradeNo string) error {
    r._alipayTradeNo = _alipayTradeNo
    r.Set("alipay_trade_no", _alipayTradeNo)
    return nil
}

// AlipayTradeNo Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAlipayTradeNo() string {
    return r._alipayTradeNo
}
// OrderType Setter
// 订单类型0 代购 1直购 3抢票
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetOrderType() int64 {
    return r._orderType
}
// AlipayAccount Setter
// 支付宝账号
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetAlipayAccount(_alipayAccount string) error {
    r._alipayAccount = _alipayAccount
    r.Set("alipay_account", _alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetAlipayAccount() string {
    return r._alipayAccount
}
// ArriveDate Setter
// 到达时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetArriveDate(_arriveDate string) error {
    r._arriveDate = _arriveDate
    r.Set("arrive_date", _arriveDate)
    return nil
}

// ArriveDate Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetArriveDate() string {
    return r._arriveDate
}
// FromStationName Setter
// 出发站
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetFromStationName(_fromStationName string) error {
    r._fromStationName = _fromStationName
    r.Set("from_station_name", _fromStationName)
    return nil
}

// FromStationName Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetFromStationName() string {
    return r._fromStationName
}
// ToStationName Setter
// 到达站
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetToStationName(_toStationName string) error {
    r._toStationName = _toStationName
    r.Set("to_station_name", _toStationName)
    return nil
}

// ToStationName Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetToStationName() string {
    return r._toStationName
}
// BoardingGates Setter
// 检票口
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetBoardingGates(_boardingGates string) error {
    r._boardingGates = _boardingGates
    r.Set("boarding_gates", _boardingGates)
    return nil
}

// BoardingGates Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetBoardingGates() string {
    return r._boardingGates
}
// ExpectDeliveryTime Setter
// 送票上门预计派送时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetExpectDeliveryTime(_expectDeliveryTime string) error {
    r._expectDeliveryTime = _expectDeliveryTime
    r.Set("expect_delivery_time", _expectDeliveryTime)
    return nil
}

// ExpectDeliveryTime Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetExpectDeliveryTime() string {
    return r._expectDeliveryTime
}
// ExtendParams Setter
// 扩展字段
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetExtendParams(_extendParams string) error {
    r._extendParams = _extendParams
    r.Set("extend_params", _extendParams)
    return nil
}

// ExtendParams Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetExtendParams() string {
    return r._extendParams
}
// TicketInfoList Setter
// 票信息列表
func (r *TaobaoTrainAgentBookticketConfirmVtwoRequest) SetTicketInfoList(_ticketInfoList []OrderTicketInfo) error {
    r._ticketInfoList = _ticketInfoList
    r.Set("ticket_info_list", _ticketInfoList)
    return nil
}

// TicketInfoList Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoRequest) GetTicketInfoList() []OrderTicketInfo {
    return r._ticketInfoList
}
