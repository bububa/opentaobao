package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdModifyAPIRequest
覆盖单元下同类型定向人群 API请求
taobao.feedflow.item.crowd.modify

覆盖单元下同类型定向人群 */
type TaobaoFeedflowItemCrowdModifyAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// New
