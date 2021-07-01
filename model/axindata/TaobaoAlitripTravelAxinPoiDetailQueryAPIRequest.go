package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest
景点poi详情查询-阿信 API请求
taobao.alitrip.travel.axin.poi.detail.query

景点poi详情查询-阿信 */
type TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest struct {
	model.Params
	// poiId
	_poiId int64
}

// New
