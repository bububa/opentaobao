package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdd 信息流增加单元
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
func TaobaoFeedflowItemAdgroupAdd(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAddAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupAddAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemAdgroupAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
