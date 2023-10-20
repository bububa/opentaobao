package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiprojectupdate 更新团期
// taobao.alitrip.travel.fsc.route.api.project.update
//
// 更新团期
func Taobaoalitriptravelfscrouteapiprojectupdate(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiprojectupdateAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiprojectupdateAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiprojectupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
