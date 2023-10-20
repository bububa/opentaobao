package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupadzonebind 信息流单元内绑定资源位
// taobao.feedflow.item.adgroup.adzone.bind
//
// 信息流单元内绑定资源位
func Taobaofeedflowitemadgroupadzonebind(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupadzonebindAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupadzonebindAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupadzonebindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
