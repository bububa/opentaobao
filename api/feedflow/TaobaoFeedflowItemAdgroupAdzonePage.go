package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupadzonepage 信息流单元下查看绑定资源位
// taobao.feedflow.item.adgroup.adzone.page
//
// 信息流单元下查看绑定资源位
func Taobaofeedflowitemadgroupadzonepage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupadzonepageAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupadzonepageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupadzonepageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
