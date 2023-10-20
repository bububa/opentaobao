package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeSetmanager OPENIM群设置管理员
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
func TaobaoOpenimTribeSetmanager(clt *core.SDKClient, req *openim.TaobaoOpenimTribeSetmanagerAPIRequest, resp *openim.TaobaoOpenimTribeSetmanagerAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
