package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayWordpackageUpdate 批量更新词包
// taobao.subway.wordpackage.update
//
// 批量更新词包
func TaobaoSubwayWordpackageUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayWordpackageUpdateAPIRequest, resp *simba.TaobaoSubwayWordpackageUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
