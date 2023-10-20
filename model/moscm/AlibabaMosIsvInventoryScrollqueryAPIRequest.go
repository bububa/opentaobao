package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosisvinventoryscrollqueryAPIRequest 滚动查询库存数据 API请求
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
type AlibabamosisvinventoryscrollqueryAPIRequest struct {
	model.Params
	// 专柜ID
	_counterId string
	// 滚动查询ID号
	_scrollId string
}

// NewAlibabamosisvinventoryscrollqueryRequest 初始化AlibabamosisvinventoryscrollqueryAPIRequest对象
func NewAlibabamosisvinventoryscrollqueryRequest() *AlibabamosisvinventoryscrollqueryAPIRequest {
	return &AlibabamosisvinventoryscrollqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosisvinventoryscrollqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.isv.inventory.scrollquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosisvinventoryscrollqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosisvinventoryscrollqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCounterId is CounterId Setter
// 专柜ID
func (r *AlibabamosisvinventoryscrollqueryAPIRequest) SetCounterId(_counterId string) error {
	r._counterId = _counterId
	r.Set("counter_id", _counterId)
	return nil
}

// GetCounterId CounterId Getter
func (r AlibabamosisvinventoryscrollqueryAPIRequest) GetCounterId() string {
	return r._counterId
}

// SetScrollId is ScrollId Setter
// 滚动查询ID号
func (r *AlibabamosisvinventoryscrollqueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabamosisvinventoryscrollqueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
