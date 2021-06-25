package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品标记接口 APIRequest
alibaba.wdk.sku.feature

给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
*/
type AlibabaWdkSkuFeatureRequest struct {
    model.Params

    // SkuFeatureDo
    param   *SkuFeatureDo 

}

func NewAlibabaWdkSkuFeatureRequest() *AlibabaWdkSkuFeatureRequest{
    return &AlibabaWdkSkuFeatureRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuFeatureRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.feature"
}

func (r AlibabaWdkSkuFeatureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuFeatureRequest) SetParam(param *SkuFeatureDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuFeatureRequest) GetParam() *SkuFeatureDo {
    return r.param
}

