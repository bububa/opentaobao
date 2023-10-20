package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiprojectclose 关闭团期
// taobao.alitrip.travel.fsc.route.api.project.close
//
// 关闭团期
func Taobaoalitriptravelfscrouteapiprojectclose(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiprojectcloseAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiprojectcloseAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiprojectcloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
