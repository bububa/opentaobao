package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCreativeofflineFind 获取创意离线多日汇总报表
// taobao.subway.creativeoffline.find
//
// 获取创意离线报表
func TaobaoSubwayCreativeofflineFind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayCreativeofflineFindAPIRequest, resp *simba.TaobaoSubwayCreativeofflineFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
