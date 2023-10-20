package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiproductadd 新增线路产品基本信息
// taobao.alitrip.travel.fsc.route.api.product.add
//
// 新增线路产品基本信息
func Taobaoalitriptravelfscrouteapiproductadd(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiproductaddAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiproductaddAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiproductaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
