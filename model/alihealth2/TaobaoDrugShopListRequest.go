package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家外卖店列表 API请求
taobao.drug.shop.list

查询卖家外卖店列表
*/
type TaobaoDrugShopListAPIRequest struct {
    model.Params
    // 查询关键字
    _keywords   string
    // 店铺状态，歇业：0，营业：1，所有：-1
    _status   int64
    // 页码
    _page   int64
    // 每页条数
    _pageSize   int64
}

// 初始化TaobaoDrugShopListAPIRequest对象
func NewTaobaoDrugShopListRequest() *TaobaoDrugShopListAPIRequest{
    return &TaobaoDrugShopListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDrugShopListAPIRequest) GetApiMethodName() string {
    return "taobao.drug.shop.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDrugShopListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keywords Setter
// 查询关键字
func (r *TaobaoDrugShopListAPIRequest) SetKeywords(_keywords string) error {
    r._keywords = _keywords
    r.Set("keywords", _keywords)
    return nil
}

// Keywords Getter
func (r TaobaoDrugShopListAPIRequest) GetKeywords() string {
    return r._keywords
}
// Status Setter
// 店铺状态，歇业：0，营业：1，所有：-1
func (r *TaobaoDrugShopListAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoDrugShopListAPIRequest) GetStatus() int64 {
    return r._status
}
// Page Setter
// 页码
func (r *TaobaoDrugShopListAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoDrugShopListAPIRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页条数
func (r *TaobaoDrugShopListAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoDrugShopListAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
