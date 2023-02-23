package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentBookticketConfirmVtwoAPIRequest 火车票代理商接口——确认出票是否成功v2--增加鉴权校验 API请求
// taobao.train.agent.bookticket.confirm.vtwo
//
// 火车票代理商接口——确认出票是否成功
type TaobaoTrainAgentBookticketConfirmVtwoAPIRequest struct {
	model.Params
	// 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
	_tickets []string
	// 票信息列表
	_ticketInfoList []OrderTicketInfo
	// 出发时间
	_depDate string
	// 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
	_failMsg string
	// 12306成功出票id
	_ticket12306Id string
	// 支付宝交易流水号
	_alipayTradeNo string
	// 错误的子订单号123434,123432
	_subOrderId string
	// 支付宝账号
	_alipayAccount string
	// 到达时间
	_arriveDate string
	// 出发站
	_fromStationName string
	// 到达站
	_toStationName string
	// 检票口
	_boardingGates string
	// 送票上门预计派送时间
	_expectDeliveryTime string
	// 扩展字段
	_extendParams string
	// 代理商id
	_agentId int64
	// 主订单id
	_mainOrderId int64
	// 订单中包含的票数量
	_ticketNum int64
	// 订单类型0 代购 1直购 3抢票
	_orderType int64
	// 是否成功
	_status bool
	// 是否支持在线退改签
	_canChange bool
}

// NewTaobaoTrainAgentBookticketConfirmVtwoRequest 初始化TaobaoTrainAgentBookticketConfirmVtwoAPIRequest对象
func NewTaobaoTrainAgentBookticketConfirmVtwoRequest() *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest {
	return &TaobaoTrainAgentBookticketConfirmVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.bookticket.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTickets is Tickets Setter
// 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetTickets(_tickets []string) error {
	r._tickets = _tickets
	r.Set("tickets", _tickets)
	return nil
}

// GetTickets Tickets Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetTickets() []string {
	return r._tickets
}

// SetTicketInfoList is TicketInfoList Setter
// 票信息列表
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetTicketInfoList(_ticketInfoList []OrderTicketInfo) error {
	r._ticketInfoList = _ticketInfoList
	r.Set("ticket_info_list", _ticketInfoList)
	return nil
}

// GetTicketInfoList TicketInfoList Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetTicketInfoList() []OrderTicketInfo {
	return r._ticketInfoList
}

// SetDepDate is DepDate Setter
// 出发时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetDepDate(_depDate string) error {
	r._depDate = _depDate
	r.Set("dep_date", _depDate)
	return nil
}

// GetDepDate DepDate Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetDepDate() string {
	return r._depDate
}

// SetFailMsg is FailMsg Setter
// 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetFailMsg(_failMsg string) error {
	r._failMsg = _failMsg
	r.Set("fail_msg", _failMsg)
	return nil
}

// GetFailMsg FailMsg Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetFailMsg() string {
	return r._failMsg
}

// SetTicket12306Id is Ticket12306Id Setter
// 12306成功出票id
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetTicket12306Id(_ticket12306Id string) error {
	r._ticket12306Id = _ticket12306Id
	r.Set("ticket_12306_id", _ticket12306Id)
	return nil
}

// GetTicket12306Id Ticket12306Id Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetTicket12306Id() string {
	return r._ticket12306Id
}

// SetAlipayTradeNo is AlipayTradeNo Setter
// 支付宝交易流水号
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetAlipayTradeNo(_alipayTradeNo string) error {
	r._alipayTradeNo = _alipayTradeNo
	r.Set("alipay_trade_no", _alipayTradeNo)
	return nil
}

// GetAlipayTradeNo AlipayTradeNo Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetAlipayTradeNo() string {
	return r._alipayTradeNo
}

// SetSubOrderId is SubOrderId Setter
// 错误的子订单号123434,123432
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetSubOrderId(_subOrderId string) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetSubOrderId() string {
	return r._subOrderId
}

// SetAlipayAccount is AlipayAccount Setter
// 支付宝账号
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetAlipayAccount(_alipayAccount string) error {
	r._alipayAccount = _alipayAccount
	r.Set("alipay_account", _alipayAccount)
	return nil
}

// GetAlipayAccount AlipayAccount Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetAlipayAccount() string {
	return r._alipayAccount
}

// SetArriveDate is ArriveDate Setter
// 到达时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetArriveDate(_arriveDate string) error {
	r._arriveDate = _arriveDate
	r.Set("arrive_date", _arriveDate)
	return nil
}

// GetArriveDate ArriveDate Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetArriveDate() string {
	return r._arriveDate
}

// SetFromStationName is FromStationName Setter
// 出发站
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetFromStationName(_fromStationName string) error {
	r._fromStationName = _fromStationName
	r.Set("from_station_name", _fromStationName)
	return nil
}

// GetFromStationName FromStationName Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetFromStationName() string {
	return r._fromStationName
}

// SetToStationName is ToStationName Setter
// 到达站
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetToStationName(_toStationName string) error {
	r._toStationName = _toStationName
	r.Set("to_station_name", _toStationName)
	return nil
}

// GetToStationName ToStationName Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetToStationName() string {
	return r._toStationName
}

// SetBoardingGates is BoardingGates Setter
// 检票口
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetBoardingGates(_boardingGates string) error {
	r._boardingGates = _boardingGates
	r.Set("boarding_gates", _boardingGates)
	return nil
}

// GetBoardingGates BoardingGates Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetBoardingGates() string {
	return r._boardingGates
}

// SetExpectDeliveryTime is ExpectDeliveryTime Setter
// 送票上门预计派送时间
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetExpectDeliveryTime(_expectDeliveryTime string) error {
	r._expectDeliveryTime = _expectDeliveryTime
	r.Set("expect_delivery_time", _expectDeliveryTime)
	return nil
}

// GetExpectDeliveryTime ExpectDeliveryTime Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetExpectDeliveryTime() string {
	return r._expectDeliveryTime
}

// SetExtendParams is ExtendParams Setter
// 扩展字段
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetTicketNum is TicketNum Setter
// 订单中包含的票数量
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetTicketNum(_ticketNum int64) error {
	r._ticketNum = _ticketNum
	r.Set("ticket_num", _ticketNum)
	return nil
}

// GetTicketNum TicketNum Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetTicketNum() int64 {
	return r._ticketNum
}

// SetOrderType is OrderType Setter
// 订单类型0 代购 1直购 3抢票
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// SetStatus is Status Setter
// 是否成功
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetStatus(_status bool) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetStatus() bool {
	return r._status
}

// SetCanChange is CanChange Setter
// 是否支持在线退改签
func (r *TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) SetCanChange(_canChange bool) error {
	r._canChange = _canChange
	r.Set("can_change", _canChange)
	return nil
}

// GetCanChange CanChange Getter
func (r TaobaoTrainAgentBookticketConfirmVtwoAPIRequest) GetCanChange() bool {
	return r._canChange
}
