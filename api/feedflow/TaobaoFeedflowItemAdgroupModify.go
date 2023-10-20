package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupModify 信息流单元修改
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
func TaobaoFeedflowItemAdgroupModify(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupModifyAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
