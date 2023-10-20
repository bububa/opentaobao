package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdd 信息流增加单元
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
func TaobaoFeedflowItemAdgroupAdd(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAddAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
