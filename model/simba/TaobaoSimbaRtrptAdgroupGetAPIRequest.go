package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRtrptAdgroupGetAPIRequest
获取推广组实时报表数据 API请求
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据 */
type TaobaoSimbaRtrptAdgroupGetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 推广计划id
	_campaignId int64
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 每页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// New
