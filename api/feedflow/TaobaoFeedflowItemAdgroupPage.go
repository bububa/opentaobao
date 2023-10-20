package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupPage 查询单元列表
// taobao.feedflow.item.adgroup.page
//
// 通过计划id查询单元信息
func TaobaoFeedflowItemAdgroupPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupPageAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
