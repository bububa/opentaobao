package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdModify 覆盖单元下同类型定向人群
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
func TaobaoFeedflowItemCrowdModify(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdModifyAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
