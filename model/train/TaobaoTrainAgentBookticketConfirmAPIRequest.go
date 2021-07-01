package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentBookticketConfirmAPIRequest
火车票代理商接口——确认出票是否成功 API请求
taobao.train.agent.bookticket.confirm

火车票代理商接口——确认出票是否成功 */
type TaobaoTrainAgentBookticketConfirmAPIRequest struct {
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
	// 火车票子订单id;单价;坐席;座次号;车次;乘车人姓名;证件类型;证件号码;保单号;保单价格 (座次号格式:坐席_车厢号_座位号，eg1:硬卧_09_03号下铺 eg2:硬座_02_03c ,注意:票价、保险价格必须到分)
	_tickets []string
	// 订单中包含的票数量
	_ticketNum int64
	// 出发时间
	_depDate string
	// 12306成功出票id
	_ticket12306Id string
	// 记录失败原因，传数字，1:票已售完,2:票价变动,3:乘车人已购相同车票,4:出票超时,5:乘车人证件未通过铁路局核验,6:发车时间变动,7:车次信息变更,8:12306故障,9:学生票信息有误,10:身份冒用,11:被限制高消费,12:坐票已售完,13:行程冲突,14:预售期变更,15:用户12306账号登录失败,16:12306账号存在未支付订单,17:用户常旅客已满,18:乘客信息有误,19:非法席别,20:车次停运,21:session登陆失败,22:账户已在其他地方登陆,23:帐号手机未核验,24:取消订单次数达到上限,25:帐号持有人身份未核验,26:邮寄地址无法保证及时送达,27:无法满足用户定制需求,28:您主动要求取消,0:未知原因
	_failMsg string
	// 支付宝交易流水号
	_alipayTradeNo string
	// 订单类型 0 代购 1直购 3抢票
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
}

// New
