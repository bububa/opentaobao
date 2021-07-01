package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJdpUserAddAPIRequest
添加数据推送用户 API请求
taobao.jushita.jdp.user.add

提供给接入数据推送的应用添加数据推送服务的用户 */
type TaobaoJushitaJdpUserAddAPIRequest struct {
	model.Params
	// RDS实例名称
	_rdsName string
	// 推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。
	_historyDays int64
}

// New
