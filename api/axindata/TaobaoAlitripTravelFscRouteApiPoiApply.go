package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapipoiapply 线路供应商提交新增POI申请
// taobao.alitrip.travel.fsc.route.api.poi.apply
//
// 线路供应商提交新增POI申请
func Taobaoalitriptravelfscrouteapipoiapply(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapipoiapplyAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapipoiapplyAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapipoiapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
