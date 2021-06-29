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
    page   int64
    // 外部档期code
    outerPromotionCode   string
    // TOB店仓编码
    ouCode   string
    // 页码大小
    pageSize   int64
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
func (r *AlibabaPricePromotionActivityQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetPage() int64 {
    return r.page
}
// OuterPromotionCode Setter
// 外部档期code
func (r *AlibabaPricePromotionActivityQueryRequest) SetOuterPromotionCode(outerPromotionCode string) error {
    r.outerPromotionCode = outerPromotionCode
    r.Set("outer_promotion_code", outerPromotionCode)
    return nil
}

// OuterPromotionCode Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetOuterPromotionCode() string {
    return r.outerPromotionCode
}
// OuCode Setter
// TOB店仓编码
func (r *AlibabaPricePromotionActivityQueryRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetOuCode() string {
    return r.ouCode
}
// PageSize Setter
// 页码大小
func (r *AlibabaPricePromotionActivityQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaPricePromotionActivityQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
