package moscm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosIsvInventoryScrollqueryAPIRequest 滚动查询库存数据 API请求
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
type AlibabaMosIsvInventoryScrollqueryAPIRequest struct {
	model.Params
	// 专柜ID
	_counterId string
	// 滚动查询ID号
	_scrollId string
}

// NewAlibabaMosIsvInventoryScrollqueryRequest 初始化AlibabaMosIsvInventoryScrollqueryAPIRequest对象
func NewAlibabaMosIsvInventoryScrollqueryRequest() *AlibabaMosIsvInventoryScrollqueryAPIRequest {
	return &AlibabaMosIsvInventoryScrollqueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosIsvInventoryScrollqueryAPIRequest) Reset() {
	r._counterId = ""
	r._scrollId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.isv.inventory.scrollquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCounterId is CounterId Setter
// 专柜ID
func (r *AlibabaMosIsvInventoryScrollqueryAPIRequest) SetCounterId(_counterId string) error {
	r._counterId = _counterId
	r.Set("counter_id", _counterId)
	return nil
}

// GetCounterId CounterId Getter
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetCounterId() string {
	return r._counterId
}

// SetScrollId is ScrollId Setter
// 滚动查询ID号
func (r *AlibabaMosIsvInventoryScrollqueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetScrollId() string {
	return r._scrollId
}

var poolAlibabaMosIsvInventoryScrollqueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosIsvInventoryScrollqueryRequest()
	},
}

// GetAlibabaMosIsvInventoryScrollqueryRequest 从 sync.Pool 获取 AlibabaMosIsvInventoryScrollqueryAPIRequest
func GetAlibabaMosIsvInventoryScrollqueryAPIRequest() *AlibabaMosIsvInventoryScrollqueryAPIRequest {
	return poolAlibabaMosIsvInventoryScrollqueryAPIRequest.Get().(*AlibabaMosIsvInventoryScrollqueryAPIRequest)
}

// ReleaseAlibabaMosIsvInventoryScrollqueryAPIRequest 将 AlibabaMosIsvInventoryScrollqueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosIsvInventoryScrollqueryAPIRequest(v *AlibabaMosIsvInventoryScrollqueryAPIRequest) {
	v.Reset()
	poolAlibabaMosIsvInventoryScrollqueryAPIRequest.Put(v)
}
