package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcreativerptdailylist 创意分日数据查询
// taobao.feedflow.item.creative.rptdailylist
//
// 创意分日数据查询
func Taobaofeedflowitemcreativerptdailylist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcreativerptdailylistAPIRequest, session string) (*feedflow.TaobaofeedflowitemcreativerptdailylistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcreativerptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
