package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU类目树获取接口 API请求
alibaba.icbu.category.get.new

获取商品发布类目
*/
type AlibabaIcbuCategoryGetNewRequest struct {
    model.Params
    // 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
    _catId   int64
}

// 初始化AlibabaIcbuCategoryGetNewRequest对象
func NewAlibabaIcbuCategoryGetNewRequest() *AlibabaIcbuCategoryGetNewRequest{
    return &AlibabaIcbuCategoryGetNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryGetNewRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.get.new"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryGetNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaIcbuCategoryGetNewRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategoryGetNewRequest) GetCatId() int64 {
    return r._catId
}
