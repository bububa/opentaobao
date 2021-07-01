package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeActivityQueryAPIRequest
查询尖货活动信息 API请求
taobao.opentrade.activity.query

尖货交易活动信息配置，查询尖货活动信息 */
type TaobaoOpentradeActivityQueryAPIRequest struct {
	model.Params
	// 活动结束时间
	_endTime string
	// 活动名称
	_activityName string
	// 分页大小
	_pageSize int64
	// 分页序号
	_pageIndex int64
}

// New
