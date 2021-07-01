package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest
查询客服在线状态 API请求
taobao.qianniu.cloudkefu.onlinestatuslog.get

按天查询客服账号的在线状态记录。如：登录，下线，挂起等
有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询 */
type TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest struct {
	model.Params
	// 子帐号列表，最多10个
	_accountIds []int64
	// 查询开始日期，只有日期有效，时间忽略
	_startDate string
	// 查询结束日期，只有日期有效，时间忽略
	_endDate string
}

// New
