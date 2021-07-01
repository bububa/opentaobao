package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemTargetValidlistAPIRequest
获取有权限的定向列表 API请求
taobao.feedflow.item.target.validlist

获取有权限的定向列表 */
type TaobaoFeedflowItemTargetValidlistAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// New
