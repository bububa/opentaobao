package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikeEserviceScheduleGetAPIRequest
客服排班信息查询接口 API请求
taobao.weike.eservice.schedule.get

客服排班信息查询接口 */
type TaobaoWeikeEserviceScheduleGetAPIRequest struct {
	model.Params
	// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
	_orderId int64
	// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
	_sellerNick string
	// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
	_spNick string
	// 起始日期，起始日期和结束日期跨度不能超过31天
	_startDate string
	// 结束日期，起始日期和结束日期跨度不能超过31天
	_endDate string
}

// New
