package gameact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// TaobaoDeActivityInfoGet 获取活动信息
// taobao.de.activity.info.get
//
// 根据appKey和活动id获取活动
func TaobaoDeActivityInfoGet(ctx context.Context, clt *core.SDKClient, req *gameact.TaobaoDeActivityInfoGetAPIRequest, resp *gameact.TaobaoDeActivityInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
