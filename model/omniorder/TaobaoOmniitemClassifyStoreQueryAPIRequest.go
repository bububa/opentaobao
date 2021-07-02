package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemClassifyStoreQueryAPIRequest 根据门店查分类信息 API请求
// taobao.omniitem.classify.store.query
//
// 根据门店查分类信息
type TaobaoOmniitemClassifyStoreQueryAPIRequest struct {
	model.Params
	// 商户的门店ID
	_storeId int64
	// 页码
	_pageNum int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoOmniitemClassifyStoreQueryRequest 初始化TaobaoOmniitemClassifyStoreQueryAPIRequest对象
func NewTaobaoOmniitemClassifyStoreQueryRequest() *TaobaoOmniitemClassifyStoreQueryAPIRequest {
	return &TaobaoOmniitemClassifyStoreQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyStoreQueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.classify.store.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyStoreQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreId is StoreId Setter
// 商户的门店ID
func (r *TaobaoOmniitemClassifyStoreQueryAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniitemClassifyStoreQueryAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetPageNum is PageNum Setter
// 页码
func (r *TaobaoOmniitemClassifyStoreQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoOmniitemClassifyStoreQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoOmniitemClassifyStoreQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOmniitemClassifyStoreQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
