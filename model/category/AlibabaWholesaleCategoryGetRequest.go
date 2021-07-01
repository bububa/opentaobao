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
type AlibabaWholesaleCategoryGetAPIRequest struct {
    model.Params
}

// 初始化AlibabaWholesaleCategoryGetAPIRequest对象
func NewAlibabaWholesaleCategoryGetRequest() *AlibabaWholesaleCategoryGetAPIRequest{
    return &AlibabaWholesaleCategoryGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleCategoryGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wholesale.category.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleCategoryGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
