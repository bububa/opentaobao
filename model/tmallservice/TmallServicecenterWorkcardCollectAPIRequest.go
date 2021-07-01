package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardCollectAPIRequest
工单揽件 API请求
tmall.servicecenter.workcard.collect

服务工单揽件接口 */
type TmallServicecenterWorkcardCollectAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 买家id
	_buyerId int64
	// 扩展信息
	_attributes string
}

// New
