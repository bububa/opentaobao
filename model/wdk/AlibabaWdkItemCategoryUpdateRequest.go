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
type AlibabaWdkItemCategoryUpdateAPIRequest struct {
    model.Params
    // 入参
    _bean   *UpdateCategoryRequestBean
}

// 初始化AlibabaWdkItemCategoryUpdateAPIRequest对象
func NewAlibabaWdkItemCategoryUpdateRequest() *AlibabaWdkItemCategoryUpdateAPIRequest{
    return &AlibabaWdkItemCategoryUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.category.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bean Setter
// 入参
func (r *AlibabaWdkItemCategoryUpdateAPIRequest) SetBean(_bean *UpdateCategoryRequestBean) error {
    r._bean = _bean
    r.Set("bean", _bean)
    return nil
}

// Bean Getter
func (r AlibabaWdkItemCategoryUpdateAPIRequest) GetBean() *UpdateCategoryRequestBean {
    return r._bean
}
