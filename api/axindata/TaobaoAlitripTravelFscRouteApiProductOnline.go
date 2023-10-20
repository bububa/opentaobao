package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiproductonline 产品上线
// taobao.alitrip.travel.fsc.route.api.product.online
//
// 产品上线
func Taobaoalitriptravelfscrouteapiproductonline(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiproductonlineAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiproductonlineAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiproductonlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
