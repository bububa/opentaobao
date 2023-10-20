package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiprojectinventoryupdate 更新团期库存
// taobao.alitrip.travel.fsc.route.api.project.inventory.update
//
// 更新团期库存
func Taobaoalitriptravelfscrouteapiprojectinventoryupdate(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiprojectinventoryupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
