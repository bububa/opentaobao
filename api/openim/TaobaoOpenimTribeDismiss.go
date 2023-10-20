package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeDismiss OPENIM群解散
// taobao.openim.tribe.dismiss
//
// OPENIM群解散
func TaobaoOpenimTribeDismiss(clt *core.SDKClient, req *openim.TaobaoOpenimTribeDismissAPIRequest, resp *openim.TaobaoOpenimTribeDismissAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
