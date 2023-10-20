package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiproductoffline 产品下线
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
func Taobaoalitriptravelfscrouteapiproductoffline(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiproductofflineAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiproductofflineAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiproductofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
