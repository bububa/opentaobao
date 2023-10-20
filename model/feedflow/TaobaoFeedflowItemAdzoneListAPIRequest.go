package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdzoneListAPIRequest 批量查询可用广告位列表 API请求
// taobao.feedflow.item.adzone.list
//
// 批量查询可用广告位列表
type TaobaoFeedflowItemAdzoneListAPIRequest struct {
	model.Params
	// 广告位查询条件
	_adzoneQuery *AdzoneQueryDto
}

// NewTaobaoFeedflowItemAdzoneListRequest 初始化TaobaoFeedflowItemAdzoneListAPIRequest对象
func NewTaobaoFeedflowItemAdzoneListRequest() *TaobaoFeedflowItemAdzoneListAPIRequest {
	return &TaobaoFeedflowItemAdzoneListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdzoneListAPIRequest) Reset() {
	r._adzoneQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneQuery is AdzoneQuery Setter
// 广告位查询条件
func (r *TaobaoFeedflowItemAdzoneListAPIRequest) SetAdzoneQuery(_adzoneQuery *AdzoneQueryDto) error {
	r._adzoneQuery = _adzoneQuery
	r.Set("adzone_query", _adzoneQuery)
	return nil
}

// GetAdzoneQuery AdzoneQuery Getter
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetAdzoneQuery() *AdzoneQueryDto {
	return r._adzoneQuery
}

var poolTaobaoFeedflowItemAdzoneListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdzoneListRequest()
	},
}

// GetTaobaoFeedflowItemAdzoneListRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdzoneListAPIRequest
func GetTaobaoFeedflowItemAdzoneListAPIRequest() *TaobaoFeedflowItemAdzoneListAPIRequest {
	return poolTaobaoFeedflowItemAdzoneListAPIRequest.Get().(*TaobaoFeedflowItemAdzoneListAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdzoneListAPIRequest 将 TaobaoFeedflowItemAdzoneListAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdzoneListAPIRequest(v *TaobaoFeedflowItemAdzoneListAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdzoneListAPIRequest.Put(v)
}
