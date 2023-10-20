package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiprojectadd 新增团期
// taobao.alitrip.travel.fsc.route.api.project.add
//
// 新增团期
func Taobaoalitriptravelfscrouteapiprojectadd(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiprojectaddAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiprojectaddAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiprojectaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
