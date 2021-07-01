package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptTargetingtageffectGetAPIRequest
获取定向效果报表数据 API请求
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据 */
type TaobaoSimbaRptTargetingtageffectGetAPIRequest struct {
	model.Params
	// 被操作者昵称
	_nick string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 起始时间
	_startTime string
	// 终止时间 ,必须小于今天
	_endTime string
	// 页面大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// New
