package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayMarshLandRptGet 获取捡漏词包分时报表数据
// taobao.subway.marsh.land.rpt.get
//
// 获取捡漏词包分时报表数据
func TaobaoSubwayMarshLandRptGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayMarshLandRptGetAPIRequest, resp *simba.TaobaoSubwayMarshLandRptGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
