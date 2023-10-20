package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowdrptdailylist 定向分日数据查询
// taobao.feedflow.item.crowd.rptdailylist
//
// 定向分日数据查询
func Taobaofeedflowitemcrowdrptdailylist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowdrptdailylistAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowdrptdailylistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowdrptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
