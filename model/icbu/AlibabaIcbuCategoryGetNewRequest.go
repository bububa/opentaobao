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
type AlibabaIcbuCategoryGetNewAPIRequest struct {
    model.Params
    // 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
    _catId   int64
}

// 初始化AlibabaIcbuCategoryGetNewAPIRequest对象
func NewAlibabaIcbuCategoryGetNewRequest() *AlibabaIcbuCategoryGetNewAPIRequest{
    return &AlibabaIcbuCategoryGetNewAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.get.new"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaIcbuCategoryGetNewAPIRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetCatId() int64 {
    return r._catId
}
