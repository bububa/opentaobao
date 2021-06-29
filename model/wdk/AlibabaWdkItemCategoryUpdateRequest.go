package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改类目 API请求
alibaba.wdk.item.category.update

修改类目
*/
type AlibabaWdkItemCategoryUpdateRequest struct {
    model.Params
    // 入参
    bean   *UpdateCategoryRequestBean
}

// 初始化AlibabaWdkItemCategoryUpdateRequest对象
func NewAlibabaWdkItemCategoryUpdateRequest() *AlibabaWdkItemCategoryUpdateRequest{
    return &AlibabaWdkItemCategoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCategoryUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.category.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCategoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bean Setter
// 入参
func (r *AlibabaWdkItemCategoryUpdateRequest) SetBean(bean *UpdateCategoryRequestBean) error {
    r.bean = bean
    r.Set("bean", bean)
    return nil
}

// Bean Getter
func (r AlibabaWdkItemCategoryUpdateRequest) GetBean() *UpdateCategoryRequestBean {
    return r.bean
}
