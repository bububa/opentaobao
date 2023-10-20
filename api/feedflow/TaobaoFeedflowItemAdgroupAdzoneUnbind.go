package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupadzoneunbind 信息流单元内解绑资源位
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
func Taobaofeedflowitemadgroupadzoneunbind(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupadzoneunbindAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupadzoneunbindAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupadzoneunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
