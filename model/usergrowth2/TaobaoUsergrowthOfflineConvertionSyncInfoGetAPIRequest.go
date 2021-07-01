package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest
查询 用增线下业务  转化数据是否同步完成 API请求
taobao.usergrowth.offline.convertion.sync.info.get

为手淘线下合作的渠道，提供对外查询数据是否更新完毕接口 */
type TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIRequest struct {
	model.Params
	// 入参
	_query *OfflineConvertionSyncInfoQuery
}

// New
