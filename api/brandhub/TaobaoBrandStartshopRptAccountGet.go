package brandhub

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// TaobaoBrandStartshopRptAccountGet 明星店铺账户报表数据查询
// taobao.brand.startshop.rpt.account.get
//
// 获取明星店铺广告主账户整体报表数据，只能查询近90天内的数据，包括展现量，点击量等
func TaobaoBrandStartshopRptAccountGet(ctx context.Context, clt *core.SDKClient, req *brandhub.TaobaoBrandStartshopRptAccountGetAPIRequest, resp *brandhub.TaobaoBrandStartshopRptAccountGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
