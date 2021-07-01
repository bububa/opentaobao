package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayAutomatchRptGetAPIRequest
查询流量智选天级报告 API请求
taobao.subway.automatch.rpt.get

查询流量智选天级报告 */
type TaobaoSubwayAutomatchRptGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 起始日期
	_startDate string
	// 终止日期
	_endDate string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
}

// New
