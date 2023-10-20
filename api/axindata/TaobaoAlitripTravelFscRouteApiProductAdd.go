package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductAdd 新增线路产品基本信息
// taobao.alitrip.travel.fsc.route.api.product.add
//
// 新增线路产品基本信息
func TaobaoAlitripTravelFscRouteApiProductAdd(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductAddAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProductAddAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProductAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
