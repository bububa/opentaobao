package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目获取接口 API请求
alibaba.wdk.sku.category.query

商家类目获取接口
*/
type AlibabaWdkSkuCategoryQueryRequest struct {
    model.Params
    // 查询类目请模型
    param   *CategoryDo
}

// 初始化AlibabaWdkSkuCategoryQueryRequest对象
func NewAlibabaWdkSkuCategoryQueryRequest() *AlibabaWdkSkuCategoryQueryRequest{
    return &AlibabaWdkSkuCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询类目请模型
func (r *AlibabaWdkSkuCategoryQueryRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuCategoryQueryRequest) GetParam() *CategoryDo {
    return r.param
}
