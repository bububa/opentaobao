package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

/* TaobaoFeedflowItemCrowdPage
分页查询单品单元下人群列表
taobao.feedflow.item.crowd.page

分页查询单品单元下人群列表 */
func TaobaoFeedflowItemCrowdPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdPageAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCrowdPageAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemCrowdPageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
