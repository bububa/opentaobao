package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdDelete 删除单品人群
// taobao.feedflow.item.crowd.delete
//
// 删除单品人群
func TaobaoFeedflowItemCrowdDelete(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
