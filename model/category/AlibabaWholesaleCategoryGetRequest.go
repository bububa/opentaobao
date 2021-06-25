package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目信息 APIRequest
alibaba.wholesale.category.get

获取类目信息
*/
type AlibabaWholesaleCategoryGetRequest struct {
    model.Params

}

func NewAlibabaWholesaleCategoryGetRequest() *AlibabaWholesaleCategoryGetRequest{
    return &AlibabaWholesaleCategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWholesaleCategoryGetRequest) GetApiMethodName() string {
    return "alibaba.wholesale.category.get"
}

func (r AlibabaWholesaleCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


