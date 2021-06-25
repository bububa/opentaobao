package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目删除接口 APIRequest
alibaba.wdk.sku.category.delete

商家类目删除接口
*/
type AlibabaWdkSkuCategoryDeleteRequest struct {
    model.Params

    // 类目删除请求模型
    param   *CategoryDo 

}

func NewAlibabaWdkSkuCategoryDeleteRequest() *AlibabaWdkSkuCategoryDeleteRequest{
    return &AlibabaWdkSkuCategoryDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCategoryDeleteRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.delete"
}

func (r AlibabaWdkSkuCategoryDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCategoryDeleteRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuCategoryDeleteRequest) GetParam() *CategoryDo {
    return r.param
}

