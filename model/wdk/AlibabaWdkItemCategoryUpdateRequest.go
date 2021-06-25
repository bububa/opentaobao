package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改类目 APIRequest
alibaba.wdk.item.category.update

修改类目
*/
type AlibabaWdkItemCategoryUpdateRequest struct {
    model.Params

    // 入参
    bean   *UpdateCategoryRequestBean 

}

func NewAlibabaWdkItemCategoryUpdateRequest() *AlibabaWdkItemCategoryUpdateRequest{
    return &AlibabaWdkItemCategoryUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemCategoryUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.category.update"
}

func (r AlibabaWdkItemCategoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemCategoryUpdateRequest) SetBean(bean *UpdateCategoryRequestBean) error {
    r.bean = bean
    r.Set("bean", bean)
    return nil
}

func (r AlibabaWdkItemCategoryUpdateRequest) GetBean() *UpdateCategoryRequestBean {
    return r.bean
}

