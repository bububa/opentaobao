package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiBusinessAreaGet 获取业务区域
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
func TaobaoAlitripTravelFscRouteApiBusinessAreaGet(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
