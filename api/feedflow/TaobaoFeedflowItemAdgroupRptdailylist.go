package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgrouprptdailylist 推广单元分日数据查询
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
func Taobaofeedflowitemadgrouprptdailylist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgrouprptdailylistAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgrouprptdailylistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgrouprptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
