package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目获取接口 APIRequest
alibaba.wdk.sku.category.query

商家类目获取接口
*/
type AlibabaWdkSkuCategoryQueryRequest struct {
    model.Params

    // 查询类目请模型
    param   *CategoryDo 

}

func NewAlibabaWdkSkuCategoryQueryRequest() *AlibabaWdkSkuCategoryQueryRequest{
    return &AlibabaWdkSkuCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.query"
}

func (r AlibabaWdkSkuCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCategoryQueryRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuCategoryQueryRequest) GetParam() *CategoryDo {
    return r.param
}

