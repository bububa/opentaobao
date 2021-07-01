package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelPoiSearchAPIRequest
POI信息查询 API请求
alitrip.travel.poi.search

POI信息查询，用于商品更新使用 */
type AlitripTravelPoiSearchAPIRequest struct {
	model.Params
	// POI信息ID；
	_id int64
	// POI信息名称
	_name string
	// POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他
	_type int64
}

// New
