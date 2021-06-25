package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU类目树获取接口 APIRequest
alibaba.icbu.category.get.new

获取商品发布类目
*/
type AlibabaIcbuCategoryGetNewRequest struct {
    model.Params

    // 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
    catId   int64 

}

func NewAlibabaIcbuCategoryGetNewRequest() *AlibabaIcbuCategoryGetNewRequest{
    return &AlibabaIcbuCategoryGetNewRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuCategoryGetNewRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.get.new"
}

func (r AlibabaIcbuCategoryGetNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuCategoryGetNewRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r AlibabaIcbuCategoryGetNewRequest) GetCatId() int64 {
    return r.catId
}

