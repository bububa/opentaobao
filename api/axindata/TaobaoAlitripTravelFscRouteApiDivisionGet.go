package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapidivisionget 获取国家城市信息
// taobao.alitrip.travel.fsc.route.api.division.get
//
// 获取国家城市信息
func Taobaoalitriptravelfscrouteapidivisionget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapidivisiongetAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapidivisiongetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapidivisiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
