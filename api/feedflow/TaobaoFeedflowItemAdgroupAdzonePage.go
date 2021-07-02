package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdzonePage 信息流单元下查看绑定资源位
// taobao.feedflow.item.adgroup.adzone.page
//
// 信息流单元下查看绑定资源位
func TaobaoFeedflowItemAdgroupAdzonePage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzonePageAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupAdzonePageAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemAdgroupAdzonePageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
