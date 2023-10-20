package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupCreativeAddBind 信息流新增并且绑定创意
// taobao.feedflow.item.adgroup.creative.add.bind
//
// 信息流新增并且绑定创意
func TaobaoFeedflowItemAdgroupCreativeAddBind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
