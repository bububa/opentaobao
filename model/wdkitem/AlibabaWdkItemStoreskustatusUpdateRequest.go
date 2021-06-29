package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品状态 API请求
alibaba.wdk.item.storeskustatus.update

五道口商品 修改门店商品状态
*/
type AlibabaWdkItemStoreskustatusUpdateRequest struct {
    model.Params
    // bean
    bean   *UpdateStoreSkuLifeStatusRequestBean
}

// 初始化AlibabaWdkItemStoreskustatusUpdateRequest对象
func NewAlibabaWdkItemStoreskustatusUpdateRequest() *AlibabaWdkItemStoreskustatusUpdateRequest{
    return &AlibabaWdkItemStoreskustatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskustatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storeskustatus.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskustatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bean Setter
// bean
func (r *AlibabaWdkItemStoreskustatusUpdateRequest) SetBean(bean *UpdateStoreSkuLifeStatusRequestBean) error {
    r.bean = bean
    r.Set("bean", bean)
    return nil
}

// Bean Getter
func (r AlibabaWdkItemStoreskustatusUpdateRequest) GetBean() *UpdateStoreSkuLifeStatusRequestBean {
    return r.bean
}
