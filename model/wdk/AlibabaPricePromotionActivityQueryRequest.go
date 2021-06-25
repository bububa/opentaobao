package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询盒马帮档期活动详情 APIRequest
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

func NewAlibabaPricePromotionActivityQueryRequest() *AlibabaPricePromotionActivityQueryRequest{
    return &AlibabaPricePromotionActivityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPricePromotionActivityQueryRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.activity.query"
}

func (r AlibabaPricePromotionActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPricePromotionActivityQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaPricePromotionActivityQueryRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaPricePromotionActivityQueryRequest) SetOuterPromotionCode(outerPromotionCode string) error {
    r.outerPromotionCode = outerPromotionCode
    r.Set("outer_promotion_code", outerPromotionCode)
    return nil
}

func (r AlibabaPricePromotionActivityQueryRequest) GetOuterPromotionCode() string {
    return r.outerPromotionCode
}

func (r *AlibabaPricePromotionActivityQueryRequest) SetOuCode(ouCode string) error {
    r.ouCode = ouCode
    r.Set("ou_code", ouCode)
    return nil
}

func (r AlibabaPricePromotionActivityQueryRequest) GetOuCode() string {
    return r.ouCode
}

func (r *AlibabaPricePromotionActivityQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaPricePromotionActivityQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

