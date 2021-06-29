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
    bizBrandKey   string
    // 每页显示数
    limit   int64
    // 是否激活
    isActivate   bool
    // 价签条码
    mac   string
    // 页码
    currentPage   int64
    // 门店ID
    storeId   int64
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
func (r *TaobaoUsceslBizApSearchRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApSearchRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}
// Limit Setter
// 每页显示数
func (r *TaobaoUsceslBizApSearchRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r TaobaoUsceslBizApSearchRequest) GetLimit() int64 {
    return r.limit
}
// IsActivate Setter
// 是否激活
func (r *TaobaoUsceslBizApSearchRequest) SetIsActivate(isActivate bool) error {
    r.isActivate = isActivate
    r.Set("is_activate", isActivate)
    return nil
}

// IsActivate Getter
func (r TaobaoUsceslBizApSearchRequest) GetIsActivate() bool {
    return r.isActivate
}
// Mac Setter
// 价签条码
func (r *TaobaoUsceslBizApSearchRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r TaobaoUsceslBizApSearchRequest) GetMac() string {
    return r.mac
}
// CurrentPage Setter
// 页码
func (r *TaobaoUsceslBizApSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoUsceslBizApSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApSearchRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApSearchRequest) GetStoreId() int64 {
    return r.storeId
}
