package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest
淘系用增线下转化明细 API请求
taobao.usergrowth.offline.convertion.details.get

淘系用增增长-线下拉新：为渠道提供返回拉新转化数据接口 */
type TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineMapiQuery
}

// New
