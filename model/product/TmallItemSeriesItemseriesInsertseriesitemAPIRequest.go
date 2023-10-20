package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSeriesItemseriesInsertseriesitemAPIRequest 向系列中添加系列商品 API请求
// tmall.item.series.itemseries.insertseriesitem
//
// 向系列中添加系列商品
type TmallItemSeriesItemseriesInsertseriesitemAPIRequest struct {
	model.Params
	// 系列商品编辑入参
	_seriesItemRequest *SeriesItemRequest
}

// NewTmallItemSeriesItemseriesInsertseriesitemRequest 初始化TmallItemSeriesItemseriesInsertseriesitemAPIRequest对象
func NewTmallItemSeriesItemseriesInsertseriesitemRequest() *TmallItemSeriesItemseriesInsertseriesitemAPIRequest {
	return &TmallItemSeriesItemseriesInsertseriesitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSeriesItemseriesInsertseriesitemAPIRequest) Reset() {
	r._seriesItemRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSeriesItemseriesInsertseriesitemAPIRequest) GetApiMethodName() string {
	return "tmall.item.series.itemseries.insertseriesitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSeriesItemseriesInsertseriesitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSeriesItemseriesInsertseriesitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesItemRequest is SeriesItemRequest Setter
// 系列商品编辑入参
func (r *TmallItemSeriesItemseriesInsertseriesitemAPIRequest) SetSeriesItemRequest(_seriesItemRequest *SeriesItemRequest) error {
	r._seriesItemRequest = _seriesItemRequest
	r.Set("series_item_request", _seriesItemRequest)
	return nil
}

// GetSeriesItemRequest SeriesItemRequest Getter
func (r TmallItemSeriesItemseriesInsertseriesitemAPIRequest) GetSeriesItemRequest() *SeriesItemRequest {
	return r._seriesItemRequest
}

var poolTmallItemSeriesItemseriesInsertseriesitemAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSeriesItemseriesInsertseriesitemRequest()
	},
}

// GetTmallItemSeriesItemseriesInsertseriesitemRequest 从 sync.Pool 获取 TmallItemSeriesItemseriesInsertseriesitemAPIRequest
func GetTmallItemSeriesItemseriesInsertseriesitemAPIRequest() *TmallItemSeriesItemseriesInsertseriesitemAPIRequest {
	return poolTmallItemSeriesItemseriesInsertseriesitemAPIRequest.Get().(*TmallItemSeriesItemseriesInsertseriesitemAPIRequest)
}

// ReleaseTmallItemSeriesItemseriesInsertseriesitemAPIRequest 将 TmallItemSeriesItemseriesInsertseriesitemAPIRequest 放入 sync.Pool
func ReleaseTmallItemSeriesItemseriesInsertseriesitemAPIRequest(v *TmallItemSeriesItemseriesInsertseriesitemAPIRequest) {
	v.Reset()
	poolTmallItemSeriesItemseriesInsertseriesitemAPIRequest.Put(v)
}
