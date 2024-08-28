package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcGroupDelete 删除指定的分组或分组下的用户
// taobao.tmc.group.delete
//
// 删除指定的分组或分组下的用户，授权消息使用
func TaobaoTmcGroupDelete(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcGroupDeleteAPIRequest, resp *tmc.TaobaoTmcGroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
