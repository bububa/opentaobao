package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupRptdailylist 推广单元分日数据查询
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
func TaobaoFeedflowItemAdgroupRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupRptdailylistAPIRequest, session string) (*feedflow.TaobaoFeedflowItemAdgroupRptdailylistAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemAdgroupRptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
