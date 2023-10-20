package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinpoidetailqueryAPIRequest 景点poi详情查询-阿信 API请求
// taobao.alitrip.travel.axin.poi.detail.query
//
// 景点poi详情查询-阿信
type TaobaoalitriptravelaxinpoidetailqueryAPIRequest struct {
	model.Params
	// poiId
	_poiId int64
}

// NewTaobaoalitriptravelaxinpoidetailqueryRequest 初始化TaobaoalitriptravelaxinpoidetailqueryAPIRequest对象
func NewTaobaoalitriptravelaxinpoidetailqueryRequest() *TaobaoalitriptravelaxinpoidetailqueryAPIRequest {
	return &TaobaoalitriptravelaxinpoidetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinpoidetailqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.poi.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinpoidetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinpoidetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoiId is PoiId Setter
// poiId
func (r *TaobaoalitriptravelaxinpoidetailqueryAPIRequest) SetPoiId(_poiId int64) error {
	r._poiId = _poiId
	r.Set("poi_id", _poiId)
	return nil
}

// GetPoiId PoiId Getter
func (r TaobaoalitriptravelaxinpoidetailqueryAPIRequest) GetPoiId() int64 {
	return r._poiId
}
