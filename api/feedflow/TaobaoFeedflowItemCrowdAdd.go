package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdAdd 单品单元下，新增定向人群
// taobao.feedflow.item.crowd.add
//
// 单品单元下，新增定向人群
func TaobaoFeedflowItemCrowdAdd(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdAddAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCrowdAddAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemCrowdAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
