package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardCompleteAPIRequest
工单完结 API请求
tmall.servicecenter.workcard.complete

工单完结 */
type TmallServicecenterWorkcardCompleteAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 完结次数
	_completeCount int64
	// 扩展信息
	_extJson string
	// 工单完结号
	_sequence int64
	// 核销地纬度
	_latitude string
	// 核销地经度
	_longitude string
}

// New
