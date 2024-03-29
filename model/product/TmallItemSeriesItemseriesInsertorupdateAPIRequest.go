package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSeriesItemseriesInsertorupdateAPIRequest 商品系列增删改接口 API请求
// tmall.item.series.itemseries.insertorupdate
//
// 商品系列增删改接口
type TmallItemSeriesItemseriesInsertorupdateAPIRequest struct {
	model.Params
	// 商品系列创建或修改请求入参
	_itemSeriesRequest *ItemSeriesRequest
}

// NewTmallItemSeriesItemseriesInsertorupdateRequest 初始化TmallItemSeriesItemseriesInsertorupdateAPIRequest对象
func NewTmallItemSeriesItemseriesInsertorupdateRequest() *TmallItemSeriesItemseriesInsertorupdateAPIRequest {
	return &TmallItemSeriesItemseriesInsertorupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSeriesItemseriesInsertorupdateAPIRequest) Reset() {
	r._itemSeriesRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSeriesItemseriesInsertorupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.series.itemseries.insertorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSeriesItemseriesInsertorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSeriesItemseriesInsertorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemSeriesRequest is ItemSeriesRequest Setter
// 商品系列创建或修改请求入参
func (r *TmallItemSeriesItemseriesInsertorupdateAPIRequest) SetItemSeriesRequest(_itemSeriesRequest *ItemSeriesRequest) error {
	r._itemSeriesRequest = _itemSeriesRequest
	r.Set("item_series_request", _itemSeriesRequest)
	return nil
}

// GetItemSeriesRequest ItemSeriesRequest Getter
func (r TmallItemSeriesItemseriesInsertorupdateAPIRequest) GetItemSeriesRequest() *ItemSeriesRequest {
	return r._itemSeriesRequest
}

var poolTmallItemSeriesItemseriesInsertorupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSeriesItemseriesInsertorupdateRequest()
	},
}

// GetTmallItemSeriesItemseriesInsertorupdateRequest 从 sync.Pool 获取 TmallItemSeriesItemseriesInsertorupdateAPIRequest
func GetTmallItemSeriesItemseriesInsertorupdateAPIRequest() *TmallItemSeriesItemseriesInsertorupdateAPIRequest {
	return poolTmallItemSeriesItemseriesInsertorupdateAPIRequest.Get().(*TmallItemSeriesItemseriesInsertorupdateAPIRequest)
}

// ReleaseTmallItemSeriesItemseriesInsertorupdateAPIRequest 将 TmallItemSeriesItemseriesInsertorupdateAPIRequest 放入 sync.Pool
func ReleaseTmallItemSeriesItemseriesInsertorupdateAPIRequest(v *TmallItemSeriesItemseriesInsertorupdateAPIRequest) {
	v.Reset()
	poolTmallItemSeriesItemseriesInsertorupdateAPIRequest.Put(v)
}
