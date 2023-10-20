package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemtargetvalidlist 获取有权限的定向列表
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
func Taobaofeedflowitemtargetvalidlist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemtargetvalidlistAPIRequest, session string) (*feedflow.TaobaofeedflowitemtargetvalidlistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemtargetvalidlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
