package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeUnsetmanager OPENIM群取消管理员
// taobao.openim.tribe.unsetmanager
//
// OPENIM群取消管理员
func TaobaoOpenimTribeUnsetmanager(clt *core.SDKClient, req *openim.TaobaoOpenimTribeUnsetmanagerAPIRequest, session string) (*openim.TaobaoOpenimTribeUnsetmanagerAPIResponse, error) {
	var resp openim.TaobaoOpenimTribeUnsetmanagerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
