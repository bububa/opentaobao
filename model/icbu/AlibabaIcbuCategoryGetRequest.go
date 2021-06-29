package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布类目获取 API请求
alibaba.icbu.category.get

获取商品发布类目
*/
type AlibabaIcbuCategoryGetRequest struct {
    model.Params
    // 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
    catId   int64
}

// 初始化AlibabaIcbuCategoryGetRequest对象
func NewAlibabaIcbuCategoryGetRequest() *AlibabaIcbuCategoryGetRequest{
    return &AlibabaIcbuCategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaIcbuCategoryGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategoryGetRequest) GetCatId() int64 {
    return r.catId
}
