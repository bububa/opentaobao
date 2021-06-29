package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目新增接口 API请求
alibaba.wdk.sku.category.add

商家类目新增接口
*/
type AlibabaWdkSkuCategoryAddRequest struct {
    model.Params
    // 类目新增请求模型
    param   *CategoryDo
}

// 初始化AlibabaWdkSkuCategoryAddRequest对象
func NewAlibabaWdkSkuCategoryAddRequest() *AlibabaWdkSkuCategoryAddRequest{
    return &AlibabaWdkSkuCategoryAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 类目新增请求模型
func (r *AlibabaWdkSkuCategoryAddRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuCategoryAddRequest) GetParam() *CategoryDo {
    return r.param
}
