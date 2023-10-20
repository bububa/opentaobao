package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiprojectopen 打开团期
// taobao.alitrip.travel.fsc.route.api.project.open
//
// 打开团期
func Taobaoalitriptravelfscrouteapiprojectopen(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiprojectopenAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiprojectopenAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiprojectopenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
