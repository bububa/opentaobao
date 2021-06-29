package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目删除接口 API请求
alibaba.wdk.sku.category.delete

商家类目删除接口
*/
type AlibabaWdkSkuCategoryDeleteRequest struct {
    model.Params
    // 类目删除请求模型
    param   *CategoryDo
}

// 初始化AlibabaWdkSkuCategoryDeleteRequest对象
func NewAlibabaWdkSkuCategoryDeleteRequest() *AlibabaWdkSkuCategoryDeleteRequest{
    return &AlibabaWdkSkuCategoryDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryDeleteRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 类目删除请求模型
func (r *AlibabaWdkSkuCategoryDeleteRequest) SetParam(param *CategoryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuCategoryDeleteRequest) GetParam() *CategoryDo {
    return r.param
}
