package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowdmodify 覆盖单元下同类型定向人群
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
func Taobaofeedflowitemcrowdmodify(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowdmodifyAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowdmodifyAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowdmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
