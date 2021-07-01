package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdDeleteAPIRequest
删除单品人群 API请求
taobao.feedflow.item.crowd.delete

删除单品人群 */
type TaobaoFeedflowItemCrowdDeleteAPIRequest struct {
	model.Params
	// 人群结构
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// New
