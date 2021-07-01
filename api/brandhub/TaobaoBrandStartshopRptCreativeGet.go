package brandhub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/brandhub"
)

/* TaobaoBrandStartshopRptCreativeGet
明星店铺创意报表数据查询
taobao.brand.startshop.rpt.creative.get

获取明星店铺广告creative分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
func TaobaoBrandStartshopRptCreativeGet(clt *core.SDKClient, req *brandhub.TaobaoBrandStartshopRptCreativeGetAPIRequest, session string) (*brandhub.TaobaoBrandStartshopRptCreativeGetAPIResponse, error) {
	var resp brandhub.TaobaoBrandStartshopRptCreativeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
