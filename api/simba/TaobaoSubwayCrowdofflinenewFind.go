package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCrowdofflinenewFind 获取人群离线多日汇总报表
// taobao.subway.crowdofflinenew.find
//
// 获取人群离线报表
func TaobaoSubwayCrowdofflinenewFind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayCrowdofflinenewFindAPIRequest, resp *simba.TaobaoSubwayCrowdofflinenewFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
