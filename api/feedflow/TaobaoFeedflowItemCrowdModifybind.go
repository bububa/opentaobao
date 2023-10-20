package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdModifybind 修改人群出价或状态
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
func TaobaoFeedflowItemCrowdModifybind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdModifybindAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdModifybindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
