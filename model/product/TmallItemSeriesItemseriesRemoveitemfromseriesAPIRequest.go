package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemseriesitemseriesremoveitemfromseriesAPIRequest 从商品系列中移除商品 API请求
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
type TmallitemseriesitemseriesremoveitemfromseriesAPIRequest struct {
	model.Params
	// 市场
	_market string
	// 商品id
	_itemId int64
	// 商品系列id
	_seriesId int64
}

// NewTmallitemseriesitemseriesremoveitemfromseriesRequest 初始化TmallitemseriesitemseriesremoveitemfromseriesAPIRequest对象
func NewTmallitemseriesitemseriesremoveitemfromseriesRequest() *TmallitemseriesitemseriesremoveitemfromseriesAPIRequest {
	return &TmallitemseriesitemseriesremoveitemfromseriesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) GetApiMethodName() string {
	return "tmall.item.series.itemseries.removeitemfromseries"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMarket is Market Setter
// 市场
func (r *TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) GetMarket() string {
	return r._market
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSeriesId is SeriesId Setter
// 商品系列id
func (r *TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) SetSeriesId(_seriesId int64) error {
	r._seriesId = _seriesId
	r.Set("series_id", _seriesId)
	return nil
}

// GetSeriesId SeriesId Getter
func (r TmallitemseriesitemseriesremoveitemfromseriesAPIRequest) GetSeriesId() int64 {
	return r._seriesId
}
