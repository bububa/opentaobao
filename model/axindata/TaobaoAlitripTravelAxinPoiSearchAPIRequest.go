package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelAxinPoiSearchAPIRequest
景点poi搜索-阿信 API请求
taobao.alitrip.travel.axin.poi.search

给阿信提供景点poi搜索 */
type TaobaoAlitripTravelAxinPoiSearchAPIRequest struct {
	model.Params
	// 搜索关键词
	_keyWord string
}

// New
