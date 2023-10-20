package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowdadd 单品单元下，新增定向人群
// taobao.feedflow.item.crowd.add
//
// 单品单元下，新增定向人群
func Taobaofeedflowitemcrowdadd(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowdaddAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowdaddAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowdaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
