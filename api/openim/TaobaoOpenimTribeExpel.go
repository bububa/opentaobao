package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeExpel OPENIM群踢出成员
// taobao.openim.tribe.expel
//
// OPENIM群踢出成员
func TaobaoOpenimTribeExpel(clt *core.SDKClient, req *openim.TaobaoOpenimTribeExpelAPIRequest, resp *openim.TaobaoOpenimTribeExpelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
