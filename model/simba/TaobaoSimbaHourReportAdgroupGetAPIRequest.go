package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaHourReportAdgroupGetAPIRequest
推广单元小时级别实时报表查询 API请求
taobao.simba.hour.report.adgroup.get

推广单元小时级别实时报表查询 */
type TaobaoSimbaHourReportAdgroupGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
	// 计划id
	_campaignId string
	// 推广单元id
	_adgroupId string
}

// New
