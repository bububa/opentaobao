package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemseriesitemseriesinsertorupdateAPIRequest 商品系列增删改接口 API请求
// tmall.item.series.itemseries.insertorupdate
//
// 商品系列增删改接口
type TmallitemseriesitemseriesinsertorupdateAPIRequest struct {
	model.Params
	// 商品系列创建或修改请求入参
	_itemSeriesRequest *ItemSeriesRequest
}

// NewTmallitemseriesitemseriesinsertorupdateRequest 初始化TmallitemseriesitemseriesinsertorupdateAPIRequest对象
func NewTmallitemseriesitemseriesinsertorupdateRequest() *TmallitemseriesitemseriesinsertorupdateAPIRequest {
	return &TmallitemseriesitemseriesinsertorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemseriesitemseriesinsertorupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.series.itemseries.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemseriesitemseriesinsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemseriesitemseriesinsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemSeriesRequest is ItemSeriesRequest Setter
// 商品系列创建或修改请求入参
func (r *TmallitemseriesitemseriesinsertorupdateAPIRequest) SetItemSeriesRequest(_itemSeriesRequest *ItemSeriesRequest) error {
	r._itemSeriesRequest = _itemSeriesRequest
	r.Set("item_series_request", _itemSeriesRequest)
	return nil
}

// GetItemSeriesRequest ItemSeriesRequest Getter
func (r TmallitemseriesitemseriesinsertorupdateAPIRequest) GetItemSeriesRequest() *ItemSeriesRequest {
	return r._itemSeriesRequest
}
