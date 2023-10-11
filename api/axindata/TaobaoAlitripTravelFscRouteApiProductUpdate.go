package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductUpdate 更新线路产品基本信息
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
func TaobaoAlitripTravelFscRouteApiProductUpdate(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
