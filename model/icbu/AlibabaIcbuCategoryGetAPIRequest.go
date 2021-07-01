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
type AlibabaIcbuCategoryGetAPIRequest struct {
    model.Params
    // 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
    _catId   int64
}

// 初始化AlibabaIcbuCategoryGetAPIRequest对象
func NewAlibabaIcbuCategoryGetRequest() *AlibabaIcbuCategoryGetAPIRequest{
    return &AlibabaIcbuCategoryGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryGetAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaIcbuCategoryGetAPIRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategoryGetAPIRequest) GetCatId() int64 {
    return r._catId
}
