package wenyuvideo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

// YoukuWenyuvideoPersionSearch 根据人物名称查询人物列表
// youku.wenyuvideo.persion.search
//
// 根据人物名称查询人物列表
func YoukuWenyuvideoPersionSearch(ctx context.Context, clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoPersionSearchAPIRequest, resp *wenyuvideo.YoukuWenyuvideoPersionSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
