package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目修改接口 API请求
alibaba.wdk.sku.category.update

商家类目修改接口
*/
type AlibabaWdkSkuCategoryUpdateRequest struct {
    model.Params
    // 更新请求模型
    _param   *CategoryDo
}

// 初始化AlibabaWdkSkuCategoryUpdateRequest对象
func NewAlibabaWdkSkuCategoryUpdateRequest() *AlibabaWdkSkuCategoryUpdateRequest{
    return &AlibabaWdkSkuCategoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 更新请求模型
func (r *AlibabaWdkSkuCategoryUpdateRequest) SetParam(_param *CategoryDo) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuCategoryUpdateRequest) GetParam() *CategoryDo {
    return r._param
}
