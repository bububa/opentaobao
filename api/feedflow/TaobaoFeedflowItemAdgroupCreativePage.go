package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupcreativepage 信息流单元下查看创意
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
func Taobaofeedflowitemadgroupcreativepage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupcreativepageAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupcreativepageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupcreativepageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
