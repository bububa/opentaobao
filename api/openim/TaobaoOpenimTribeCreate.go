package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeCreate 创建群
// taobao.openim.tribe.create
//
// 创建一个openim的群
func TaobaoOpenimTribeCreate(clt *core.SDKClient, req *openim.TaobaoOpenimTribeCreateAPIRequest, resp *openim.TaobaoOpenimTribeCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
