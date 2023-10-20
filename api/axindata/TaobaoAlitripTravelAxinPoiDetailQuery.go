package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinPoiDetailQuery 景点poi详情查询-阿信
// taobao.alitrip.travel.axin.poi.detail.query
//
// 景点poi详情查询-阿信
func TaobaoAlitripTravelAxinPoiDetailQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest, resp *axindata.TaobaoAlitripTravelAxinPoiDetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
