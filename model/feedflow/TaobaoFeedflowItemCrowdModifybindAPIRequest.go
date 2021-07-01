package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdModifybindAPIRequest
修改人群出价或状态 API请求
taobao.feedflow.item.crowd.modifybind

修改人群出价或状态 */
type TaobaoFeedflowItemCrowdModifybindAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// New
