package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeModifytribeinfo OPENIM群信息修改
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
func TaobaoOpenimTribeModifytribeinfo(clt *core.SDKClient, req *openim.TaobaoOpenimTribeModifytribeinfoAPIRequest, resp *openim.TaobaoOpenimTribeModifytribeinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
