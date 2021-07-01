package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaReportCityGetAPIRequest
获取城市维度报表 API请求
taobao.simba.report.city.get

获取城市维度报表 */
type TaobaoSimbaReportCityGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
	// 计划id
	_campaignId string
}

// New
