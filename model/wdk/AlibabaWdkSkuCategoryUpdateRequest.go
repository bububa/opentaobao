package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目修改接口 APIRequest
alibaba.wdk.sku.category.update

商家类目修改接口
*/
type AlibabaWdkSkuCategoryUpdateRequest struct {
    model.Params

    // 更新请求模型
    param   *CategoryDo 

}

func NewAlibabaWdkSkuCategoryUpdateRequest() *AlibabaWdkSkuCategoryUpdateRequest{
    return &AlibabaWdkSkuCategoryUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCategoryUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.update"
}

func (r AlibabaWdkSkuCategoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCategoryUpdateRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuCategoryUpdateRequest) GetParam() *CategoryDo {
    return r.param
}

