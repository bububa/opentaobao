package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目新增接口 APIRequest
alibaba.wdk.sku.category.add

商家类目新增接口
*/
type AlibabaWdkSkuCategoryAddRequest struct {
    model.Params

    // 类目新增请求模型
    param   *CategoryDo 

}

func NewAlibabaWdkSkuCategoryAddRequest() *AlibabaWdkSkuCategoryAddRequest{
    return &AlibabaWdkSkuCategoryAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCategoryAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.add"
}

func (r AlibabaWdkSkuCategoryAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCategoryAddRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuCategoryAddRequest) GetParam() *CategoryDo {
    return r.param
}

