package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeModifytribeinfo OPENIM群信息修改
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
func TaobaoOpenimTribeModifytribeinfo(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeModifytribeinfoAPIRequest, resp *openim.TaobaoOpenimTribeModifytribeinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
