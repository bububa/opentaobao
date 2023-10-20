package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest 从商品系列中移除商品 API请求
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
type TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest struct {
	model.Params
	// 市场
	_market string
	// 商品id
	_itemId int64
	// 商品系列id
	_seriesId int64
}

// NewTmallItemSeriesItemseriesRemoveitemfromseriesRequest 初始化TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest对象
func NewTmallItemSeriesItemseriesRemoveitemfromseriesRequest() *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest {
	return &TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) Reset() {
	r._market = ""
	r._itemId = 0
	r._seriesId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) GetApiMethodName() string {
	return "tmall.item.series.itemseries.removeitemfromseries"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMarket is Market Setter
// 市场
func (r *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) SetMarket(_market string) error {
	r._market = _market
	r.Set("market", _market)
	return nil
}

// GetMarket Market Getter
func (r TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) GetMarket() string {
	return r._market
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSeriesId is SeriesId Setter
// 商品系列id
func (r *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) SetSeriesId(_seriesId int64) error {
	r._seriesId = _seriesId
	r.Set("series_id", _seriesId)
	return nil
}

// GetSeriesId SeriesId Getter
func (r TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) GetSeriesId() int64 {
	return r._seriesId
}

var poolTmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSeriesItemseriesRemoveitemfromseriesRequest()
	},
}

// GetTmallItemSeriesItemseriesRemoveitemfromseriesRequest 从 sync.Pool 获取 TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest
func GetTmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest() *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest {
	return poolTmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest.Get().(*TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest)
}

// ReleaseTmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest 将 TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest 放入 sync.Pool
func ReleaseTmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest(v *TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest) {
	v.Reset()
	poolTmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest.Put(v)
}
