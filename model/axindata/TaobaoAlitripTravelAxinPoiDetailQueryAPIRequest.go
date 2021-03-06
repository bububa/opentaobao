package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest 景点poi详情查询-阿信 API请求
// taobao.alitrip.travel.axin.poi.detail.query
//
// 景点poi详情查询-阿信
type TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest struct {
	model.Params
	// poiId
	_poiId int64
}

// NewTaobaoAlitripTravelAxinPoiDetailQueryRequest 初始化TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest对象
func NewTaobaoAlitripTravelAxinPoiDetailQueryRequest() *TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest {
	return &TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.poi.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPoiId is PoiId Setter
// poiId
func (r *TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest) SetPoiId(_poiId int64) error {
	r._poiId = _poiId
	r.Set("poi_id", _poiId)
	return nil
}

// GetPoiId PoiId Getter
func (r TaobaoAlitripTravelAxinPoiDetailQueryAPIRequest) GetPoiId() int64 {
	return r._poiId
}
