package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询盒马帮档期活动详情 API请求
alibaba.price.promotion.activity.query

查询盒马帮档期活动详情
*/
type AlibabaPricePromotionActivityQueryRequest struct {
    model.Params
    // 页码
    _page   int64
    // 外部档期code
    _outerPromotionCode   string
    // TOB店仓编码
    _ouCode   string
    // 页码大小
    _pageSize   int64
}

// 初始化AlibabaPricePromotionActivityQueryRequest对象
func NewAlibabaPricePromotionActivityQueryRequest() *AlibabaPricePromotionActivityQueryRequest{
    return &AlibabaPricePromotionActivityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionActivityQueryRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Page Setter
// 页码
func (r *AlibabaPricePromotionActivityQueryRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetPage() int64 {
    return r._page
}
// OuterPromotionCode Setter
// 外部档期code
func (r *AlibabaPricePromotionActivityQueryRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
    r._outerPromotionCode = _outerPromotionCode
    r.Set("outer_promotion_code", _outerPromotionCode)
    return nil
}

// OuterPromotionCode Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetOuterPromotionCode() string {
    return r._outerPromotionCode
}
// OuCode Setter
// TOB店仓编码
func (r *AlibabaPricePromotionActivityQueryRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetOuCode() string {
    return r._ouCode
}
// PageSize Setter
// 页码大小
func (r *AlibabaPricePromotionActivityQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
