package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdModifybind 修改人群出价或状态
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
func TaobaoFeedflowItemCrowdModifybind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdModifybindAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCrowdModifybindAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemCrowdModifybindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
