package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapiproductlabelget 获取线路主题
// taobao.alitrip.travel.fsc.route.api.product.label.get
//
// 获取线路主题
func Taobaoalitriptravelfscrouteapiproductlabelget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapiproductlabelgetAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapiproductlabelgetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapiproductlabelgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
