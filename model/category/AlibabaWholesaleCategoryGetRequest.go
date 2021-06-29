package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目信息 API请求
alibaba.wholesale.category.get

获取类目信息
*/
type AlibabaWholesaleCategoryGetRequest struct {
    model.Params
}

// 初始化AlibabaWholesaleCategoryGetRequest对象
func NewAlibabaWholesaleCategoryGetRequest() *AlibabaWholesaleCategoryGetRequest{
    return &AlibabaWholesaleCategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleCategoryGetRequest) GetApiMethodName() string {
    return "alibaba.wholesale.category.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
