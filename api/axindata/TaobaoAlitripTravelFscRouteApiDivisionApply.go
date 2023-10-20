package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiDivisionApply 线路供应商提交新增城市申请
// taobao.alitrip.travel.fsc.route.api.division.apply
//
// 线路供应商提交新增城市申请
func TaobaoAlitripTravelFscRouteApiDivisionApply(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiDivisionApplyAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiDivisionApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
