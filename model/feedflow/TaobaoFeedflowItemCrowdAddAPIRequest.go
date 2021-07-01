package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdAddAPIRequest
单品单元下，新增定向人群 API请求
taobao.feedflow.item.crowd.add

单品单元下，新增定向人群 */
type TaobaoFeedflowItemCrowdAddAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// New
