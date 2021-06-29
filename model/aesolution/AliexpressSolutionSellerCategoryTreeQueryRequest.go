package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.seller.category.tree.query APIRequest
aliexpress.solution.seller.category.tree.query

API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
*/
type AliexpressSolutionSellerCategoryTreeQueryRequest struct {
    model.Params

    // parent category ID. To obtain the root categories, setting the category_id = 0
    categoryId   int64 

    // filter the categories which seller does not have permission
    filterNoPermission   bool 

}

func NewAliexpressSolutionSellerCategoryTreeQueryRequest() *AliexpressSolutionSellerCategoryTreeQueryRequest{
    return &AliexpressSolutionSellerCategoryTreeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetApiMethodName() string {
    return "aliexpress.solution.seller.category.tree.query"
}

func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionSellerCategoryTreeQueryRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *AliexpressSolutionSellerCategoryTreeQueryRequest) SetFilterNoPermission(filterNoPermission bool) error {
    r.filterNoPermission = filterNoPermission
    r.Set("filter_no_permission", filterNoPermission)
    return nil
}

func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetFilterNoPermission() bool {
    return r.filterNoPermission
}

