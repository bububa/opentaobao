package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductLabelGet 获取线路主题
// taobao.alitrip.travel.fsc.route.api.product.label.get
//
// 获取线路主题
func TaobaoAlitripTravelFscRouteApiProductLabelGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductLabelGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiProductLabelGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
