package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiproductupdate 更新线路产品基本信息
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
func Taobaoalitriptravelfscrouteapiproductupdate(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiproductupdateAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiproductupdateAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiproductupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
