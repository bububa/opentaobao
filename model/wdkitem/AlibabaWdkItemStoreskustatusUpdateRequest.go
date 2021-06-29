package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店商品状态 APIRequest
alibaba.wdk.item.storeskustatus.update

五道口商品 修改门店商品状态
*/
type AlibabaWdkItemStoreskustatusUpdateRequest struct {
    model.Params

    // bean
    bean   *UpdateStoreSkuLifeStatusRequestBean 

}

func NewAlibabaWdkItemStoreskustatusUpdateRequest() *AlibabaWdkItemStoreskustatusUpdateRequest{
    return &AlibabaWdkItemStoreskustatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemStoreskustatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.storeskustatus.update"
}

func (r AlibabaWdkItemStoreskustatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemStoreskustatusUpdateRequest) SetBean(bean *UpdateStoreSkuLifeStatusRequestBean) error {
    r.bean = bean
    r.Set("bean", bean)
    return nil
}

func (r AlibabaWdkItemStoreskustatusUpdateRequest) GetBean() *UpdateStoreSkuLifeStatusRequestBean {
    return r.bean
}

