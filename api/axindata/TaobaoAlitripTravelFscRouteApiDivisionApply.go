package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapidivisionapply 线路供应商提交新增城市申请
// taobao.alitrip.travel.fsc.route.api.division.apply
//
// 线路供应商提交新增城市申请
func Taobaoalitriptravelfscrouteapidivisionapply(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapidivisionapplyAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapidivisionapplyAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapidivisionapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
