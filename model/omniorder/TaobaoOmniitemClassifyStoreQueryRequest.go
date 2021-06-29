package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据门店查分类信息 API请求
taobao.omniitem.classify.store.query

根据门店查分类信息
*/
type TaobaoOmniitemClassifyStoreQueryRequest struct {
    model.Params
    // 商户的门店ID
    _storeId   int64
    // 页码
    _pageNum   int64
    // 每页大小
    _pageSize   int64
}

// 初始化TaobaoOmniitemClassifyStoreQueryRequest对象
func NewTaobaoOmniitemClassifyStoreQueryRequest() *TaobaoOmniitemClassifyStoreQueryRequest{
    return &TaobaoOmniitemClassifyStoreQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyStoreQueryRequest) GetApiMethodName() string {
    return "taobao.omniitem.classify.store.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyStoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 商户的门店ID
func (r *TaobaoOmniitemClassifyStoreQueryRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoOmniitemClassifyStoreQueryRequest) GetStoreId() int64 {
    return r._storeId
}
// PageNum Setter
// 页码
func (r *TaobaoOmniitemClassifyStoreQueryRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoOmniitemClassifyStoreQueryRequest) GetPageNum() int64 {
    return r._pageNum
}
// PageSize Setter
// 每页大小
func (r *TaobaoOmniitemClassifyStoreQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOmniitemClassifyStoreQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
