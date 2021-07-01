package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentBookticketConfirmVtwoAPIRequest
火车票代理商接口——确认出票是否成功v2--增加鉴权校验 API请求
taobao.train.agent.bookticket.confirm.vtwo

火车票代理商接口——确认出票是否成功 */
type TaobaoTrainAgentBookticketConfirmVtwoAPIRequest struct {
	model.Params
	// 错误的子订单号123434,123432
	_subOrderId string
	// 是否支持在线退改签
	_canChange bool
	// 主订单id
	_mainOrderId int64
	// 是否成功
	_status bool
	// 代理商id
	_agentId int64
	// 火车票id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格特别注意:票价、保险价格必须到分,例如10元,输入为1000.
	_tickets []string
	// 订单中包含的票数量
	_ticketNum int64
	// 出发时间
	_depDate string
	// 12306成功出票id
	_ticket12306Id string
	// 记录失败原因，传数字，1、票已售完，出票失败全额退款，2、票价变动，出票失败全额退款，3、乘车人已购买相同车票，出票失败全额退款， 4、出票超时，出票失败全额退款，5、乘车人证件未通过铁路局审核，需到售票窗口办理， 6、发车时间变动，出票失败全额退款，7、车次信息错误，出票失败全额退款，8、12306故障,出票失败全额退款， 0、出票失败全额退款
	_failMsg string
	// 支付宝交易流水号
	_alipayTradeNo string
	// 订单类型0 代购 1直购 3抢票
	_orderType int64
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
	// 票信息列表
	_ticketInfoList []OrderTicketInfo
}

// New
