package brandhub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

// TaobaoBrandStartshopRptCreativeGet 明星店铺创意报表数据查询
// taobao.brand.startshop.rpt.creative.get
//
// 获取明星店铺广告creative分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
func TaobaoBrandStartshopRptCreativeGet(clt *core.SDKClient, req *brandhub.TaobaoBrandStartshopRptCreativeGetAPIRequest, resp *brandhub.TaobaoBrandStartshopRptCreativeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
