package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AP列表查询 API请求
taobao.uscesl.biz.ap.search

查询当前门店下登记的AP列表
*/
type TaobaoUsceslBizApSearchRequest struct {
    model.Params
    // 商家编码
    _bizBrandKey   string
    // 每页显示数
    _limit   int64
    // 是否激活
    _isActivate   bool
    // 价签条码
    _mac   string
    // 页码
    _currentPage   int64
    // 门店ID
    _storeId   int64
}

// 初始化TaobaoUsceslBizApSearchRequest对象
func NewTaobaoUsceslBizApSearchRequest() *TaobaoUsceslBizApSearchRequest{
    return &TaobaoUsceslBizApSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApSearchRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizBrandKey Setter
// 商家编码
func (r *TaobaoUsceslBizApSearchRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApSearchRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}
// Limit Setter
// 每页显示数
func (r *TaobaoUsceslBizApSearchRequest) SetLimit(_limit int64) error {
    r._limit = _limit
    r.Set("limit", _limit)
    return nil
}

// Limit Getter
func (r TaobaoUsceslBizApSearchRequest) GetLimit() int64 {
    return r._limit
}
// IsActivate Setter
// 是否激活
func (r *TaobaoUsceslBizApSearchRequest) SetIsActivate(_isActivate bool) error {
    r._isActivate = _isActivate
    r.Set("is_activate", _isActivate)
    return nil
}

// IsActivate Getter
func (r TaobaoUsceslBizApSearchRequest) GetIsActivate() bool {
    return r._isActivate
}
// Mac Setter
// 价签条码
func (r *TaobaoUsceslBizApSearchRequest) SetMac(_mac string) error {
    r._mac = _mac
    r.Set("mac", _mac)
    return nil
}

// Mac Getter
func (r TaobaoUsceslBizApSearchRequest) GetMac() string {
    return r._mac
}
// CurrentPage Setter
// 页码
func (r *TaobaoUsceslBizApSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoUsceslBizApSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApSearchRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApSearchRequest) GetStoreId() int64 {
    return r._storeId
}
