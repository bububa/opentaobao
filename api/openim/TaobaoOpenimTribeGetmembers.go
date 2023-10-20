package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeGetmembers OPENIM群成员获取
// taobao.openim.tribe.getmembers
//
// OPENIM群成员获取
func TaobaoOpenimTribeGetmembers(clt *core.SDKClient, req *openim.TaobaoOpenimTribeGetmembersAPIRequest, resp *openim.TaobaoOpenimTribeGetmembersAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
