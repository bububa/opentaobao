package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardUpdatelogisticsAPIRequest
更新物流进度 API请求
tmall.servicecenter.workcard.updatelogistics

提供给外部合作服务商的物流进度更改接口 */
type TmallServicecenterWorkcardUpdatelogisticsAPIRequest struct {
	model.Params
	// 工单号
	_workcardId int64
	// 工单操作
	_action string
	// 快递公司
	_expressCompany string
	// 快递号
	_expressCode string
}

// New
