package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinPoiSearch 景点poi搜索-阿信
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
func TaobaoAlitripTravelAxinPoiSearch(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinPoiSearchAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinPoiSearchAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinPoiSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
