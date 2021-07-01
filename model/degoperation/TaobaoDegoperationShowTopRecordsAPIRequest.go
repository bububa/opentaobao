package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationShowTopRecordsAPIRequest
活动中奖记录 API请求
taobao.degoperation.show.top.records

活动中奖记录 */
type TaobaoDegoperationShowTopRecordsAPIRequest struct {
	model.Params
	// 活动后台配置
	_degAppKey string
	// 活动后台配置
	_degEventKey string
	// 返回数
	_topN int64
}

// New
