package brandhub

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// TaobaoBrandStartshopRptAdgroupGet 明星店铺推广单元报表数据查询
// taobao.brand.startshop.rpt.adgroup.get
//
// 获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func TaobaoBrandStartshopRptAdgroupGet(ctx context.Context, clt *core.SDKClient, req *brandhub.TaobaoBrandStartshopRptAdgroupGetAPIRequest, resp *brandhub.TaobaoBrandStartshopRptAdgroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
