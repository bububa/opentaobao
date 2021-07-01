package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRtrptTargetingtagGetAPIRequest
搜索人群实时报表 API请求
taobao.simba.rtrpt.targetingtag.get

获取搜搜人群实时报表 */
type TaobaoSimbaRtrptTargetingtagGetAPIRequest struct {
	model.Params
	// 旺旺名称
	_nick string
	// 推广计划id
	_campaignId int64
	// 推广单元id
	_adgroupId int64
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
}

// New
