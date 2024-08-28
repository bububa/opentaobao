package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCreativeUpdate 修改创意
// taobao.simba.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
func TaobaoSimbaCreativeUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCreativeUpdateAPIRequest, resp *simba.TaobaoSimbaCreativeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
