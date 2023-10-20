package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowdpage 分页查询单品单元下人群列表
// taobao.feedflow.item.crowd.page
//
// 分页查询单品单元下人群列表
func Taobaofeedflowitemcrowdpage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowdpageAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowdpageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowdpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
