package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeJoin OPENIM群主动加入
// taobao.openim.tribe.join
//
// OPENIM群主动加入
func TaobaoOpenimTribeJoin(clt *core.SDKClient, req *openim.TaobaoOpenimTribeJoinAPIRequest, resp *openim.TaobaoOpenimTribeJoinAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
