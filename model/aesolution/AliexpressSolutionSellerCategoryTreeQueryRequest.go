package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.seller.category.tree.query API请求
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

// 初始化AliexpressSolutionSellerCategoryTreeQueryRequest对象
func NewAliexpressSolutionSellerCategoryTreeQueryRequest() *AliexpressSolutionSellerCategoryTreeQueryRequest{
    return &AliexpressSolutionSellerCategoryTreeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetApiMethodName() string {
    return "aliexpress.solution.seller.category.tree.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// parent category ID. To obtain the root categories, setting the category_id = 0
func (r *AliexpressSolutionSellerCategoryTreeQueryRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetCategoryId() int64 {
    return r.categoryId
}
// FilterNoPermission Setter
// filter the categories which seller does not have permission
func (r *AliexpressSolutionSellerCategoryTreeQueryRequest) SetFilterNoPermission(filterNoPermission bool) error {
    r.filterNoPermission = filterNoPermission
    r.Set("filter_no_permission", filterNoPermission)
    return nil
}

// FilterNoPermission Getter
func (r AliexpressSolutionSellerCategoryTreeQueryRequest) GetFilterNoPermission() bool {
    return r.filterNoPermission
}
