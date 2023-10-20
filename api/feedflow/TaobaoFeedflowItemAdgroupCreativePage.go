package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupCreativePage 信息流单元下查看创意
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
func TaobaoFeedflowItemAdgroupCreativePage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupCreativePageAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupCreativePageAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemAdgroupCreativePageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
