package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeGetalltribes 获取用户群列表
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
func TaobaoOpenimTribeGetalltribes(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribeGetalltribesAPIRequest, resp *openim.TaobaoOpenimTribeGetalltribesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
