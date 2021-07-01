package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest
手淘线下拉新业务 t+8转化明细数据 API请求
taobao.usergrowth.offline.convertion.details.eight.get

手淘线下拉新业务 给合作渠道返回t+8转化明细数据 */
type TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineMapiQuery
}

// New
