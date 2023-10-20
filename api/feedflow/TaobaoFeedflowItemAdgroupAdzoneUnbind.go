package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdzoneUnbind 信息流单元内解绑资源位
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
func TaobaoFeedflowItemAdgroupAdzoneUnbind(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
