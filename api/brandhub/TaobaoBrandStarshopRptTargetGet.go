package brandhub

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// TaobaoBrandStarshopRptTargetGet 明星店铺定向维度报表
// taobao.brand.starshop.rpt.target.get
//
// 获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func TaobaoBrandStarshopRptTargetGet(ctx context.Context, clt *core.SDKClient, req *brandhub.TaobaoBrandStarshopRptTargetGetAPIRequest, resp *brandhub.TaobaoBrandStarshopRptTargetGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
