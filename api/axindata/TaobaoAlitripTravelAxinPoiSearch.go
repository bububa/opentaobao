package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinPoiSearch 景点poi搜索-阿信
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
func TaobaoAlitripTravelAxinPoiSearch(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinPoiSearchAPIRequest, resp *axindata.TaobaoAlitripTravelAxinPoiSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
