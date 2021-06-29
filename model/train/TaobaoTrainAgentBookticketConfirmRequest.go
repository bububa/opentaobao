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
    _subOrderId   string
    // 是否支持在线退改签
    _canChange   bool
    // 主订单id
    _mainOrderId   int64
    // 是否成功
    _status   bool
    // 代理商id
    _agentId   int64
    // 火车票子订单id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格 (座次号格式:坐席_车厢号_座位号，eg1:硬卧_09_03号下铺 eg2:硬座_02_03c ,注意:票价、保险价格必须到分)
    _tickets   []string
    // 订单中包含的票数量
    _ticketNum   int64
    // 出发时间
    _depDate   string
    // 12306成功出票id
    _ticket12306Id   string
    // 记录失败原因，传数字，1:票已售完,2:票价变动,3:乘车人已购相同车票,4:出票超时,5:乘车人证件未通过铁路局核验,6:发车时间变动,7:车次信息变更,8:12306故障,9:学生票信息有误,10:身份冒用,11:被限制高消费,12:坐票已售完,13:行程冲突,14:预售期变更,15:用户12306账号登录失败,16:12306账号存在未支付订单,17:用户常旅客已满,18:乘客信息有误,19:非法席别,20:车次停运,21:session登陆失败,22:账户已在其他地方登陆,23:帐号手机未核验,24:取消订单次数达到上限,25:帐号持有人身份未核验,26:邮寄地址无法保证及时送达,27:无法满足用户定制需求,28:您主动要求取消,0:未知原因
    _failMsg   string
    // 支付宝交易流水号
    _alipayTradeNo   string
    // 订单类型 0 代购 1直购 3抢票
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
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetSubOrderId(_subOrderId string) error {
    r._subOrderId = _subOrderId
    r.Set("sub_order_id", _subOrderId)
    return nil
}

// SubOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetSubOrderId() string {
    return r._subOrderId
}
// CanChange Setter
// 是否支持在线退改签
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetCanChange(_canChange bool) error {
    r._canChange = _canChange
    r.Set("can_change", _canChange)
    return nil
}

// CanChange Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetCanChange() bool {
    return r._canChange
}
// MainOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// Status Setter
// 是否成功
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetStatus(_status bool) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetStatus() bool {
    return r._status
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetAgentId() int64 {
    return r._agentId
}
// Tickets Setter
// 火车票子订单id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格 (座次号格式:坐席_车厢号_座位号，eg1:硬卧_09_03号下铺 eg2:硬座_02_03c ,注意:票价、保险价格必须到分)
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetTickets(_tickets []string) error {
    r._tickets = _tickets
    r.Set("tickets", _tickets)
    return nil
}

// Tickets Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetTickets() []string {
    return r._tickets
}
// TicketNum Setter
// 订单中包含的票数量
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetTicketNum(_ticketNum int64) error {
    r._ticketNum = _ticketNum
    r.Set("ticket_num", _ticketNum)
    return nil
}

// TicketNum Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetTicketNum() int64 {
    return r._ticketNum
}
// DepDate Setter
// 出发时间
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetDepDate(_depDate string) error {
    r._depDate = _depDate
    r.Set("dep_date", _depDate)
    return nil
}

// DepDate Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetDepDate() string {
    return r._depDate
}
// Ticket12306Id Setter
// 12306成功出票id
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetTicket12306Id(_ticket12306Id string) error {
    r._ticket12306Id = _ticket12306Id
    r.Set("ticket_12306_id", _ticket12306Id)
    return nil
}

// Ticket12306Id Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetTicket12306Id() string {
    return r._ticket12306Id
}
// FailMsg Setter
// 记录失败原因，传数字，1:票已售完,2:票价变动,3:乘车人已购相同车票,4:出票超时,5:乘车人证件未通过铁路局核验,6:发车时间变动,7:车次信息变更,8:12306故障,9:学生票信息有误,10:身份冒用,11:被限制高消费,12:坐票已售完,13:行程冲突,14:预售期变更,15:用户12306账号登录失败,16:12306账号存在未支付订单,17:用户常旅客已满,18:乘客信息有误,19:非法席别,20:车次停运,21:session登陆失败,22:账户已在其他地方登陆,23:帐号手机未核验,24:取消订单次数达到上限,25:帐号持有人身份未核验,26:邮寄地址无法保证及时送达,27:无法满足用户定制需求,28:您主动要求取消,0:未知原因
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetFailMsg(_failMsg string) error {
    r._failMsg = _failMsg
    r.Set("fail_msg", _failMsg)
    return nil
}

// FailMsg Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetFailMsg() string {
    return r._failMsg
}
// AlipayTradeNo Setter
// 支付宝交易流水号
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetAlipayTradeNo(_alipayTradeNo string) error {
    r._alipayTradeNo = _alipayTradeNo
    r.Set("alipay_trade_no", _alipayTradeNo)
    return nil
}

// AlipayTradeNo Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetAlipayTradeNo() string {
    return r._alipayTradeNo
}
// OrderType Setter
// 订单类型 0 代购 1直购 3抢票
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetOrderType() int64 {
    return r._orderType
}
// AlipayAccount Setter
// 支付宝账号
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetAlipayAccount(_alipayAccount string) error {
    r._alipayAccount = _alipayAccount
    r.Set("alipay_account", _alipayAccount)
    return nil
}

// AlipayAccount Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetAlipayAccount() string {
    return r._alipayAccount
}
// ArriveDate Setter
// 到达时间
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetArriveDate(_arriveDate string) error {
    r._arriveDate = _arriveDate
    r.Set("arrive_date", _arriveDate)
    return nil
}

// ArriveDate Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetArriveDate() string {
    return r._arriveDate
}
// FromStationName Setter
// 出发站
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetFromStationName(_fromStationName string) error {
    r._fromStationName = _fromStationName
    r.Set("from_station_name", _fromStationName)
    return nil
}

// FromStationName Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetFromStationName() string {
    return r._fromStationName
}
// ToStationName Setter
// 到达站
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetToStationName(_toStationName string) error {
    r._toStationName = _toStationName
    r.Set("to_station_name", _toStationName)
    return nil
}

// ToStationName Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetToStationName() string {
    return r._toStationName
}
// BoardingGates Setter
// 检票口
func (r *TaobaoTrainAgentBookticketConfirmRequest) SetBoardingGates(_boardingGates string) error {
    r._boardingGates = _boardingGates
    r.Set("boarding_gates", _boardingGates)
    return nil
}

// BoardingGates Getter
func (r TaobaoTrainAgentBookticketConfirmRequest) GetBoardingGates() string {
    return r._boardingGates
}
