package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterIdentifytaskCreateAPIRequest
服务商创建核销单 API请求
tmall.servicecenter.identifytask.create

服务商调用该接口进行创建核销单操作 */
type TmallServicecenterIdentifytaskCreateAPIRequest struct {
	model.Params
	// 工单列表
	_workcardIds []int64
	// 是否改派
	_reassign bool
	// 服务商自定义的外部核销单id
	_outerId string
}

// New
