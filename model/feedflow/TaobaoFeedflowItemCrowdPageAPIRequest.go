package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdPageAPIRequest
分页查询单品单元下人群列表 API请求
taobao.feedflow.item.crowd.page

分页查询单品单元下人群列表 */
type TaobaoFeedflowItemCrowdPageAPIRequest struct {
	model.Params
	// 查询条件
	_crowdQuery *CrowdQueryDto
}

// New
