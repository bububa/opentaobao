package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupDelete 根据单元id删除单元
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
func TaobaoFeedflowItemAdgroupDelete(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
