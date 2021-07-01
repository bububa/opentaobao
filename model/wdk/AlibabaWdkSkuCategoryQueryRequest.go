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
type AlibabaWdkSkuCategoryQueryAPIRequest struct {
    model.Params
    // 查询类目请模型
    _param   *CategoryDO
}

// 初始化AlibabaWdkSkuCategoryQueryAPIRequest对象
func NewAlibabaWdkSkuCategoryQueryRequest() *AlibabaWdkSkuCategoryQueryAPIRequest{
    return &AlibabaWdkSkuCategoryQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.category.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询类目请模型
func (r *AlibabaWdkSkuCategoryQueryAPIRequest) SetParam(_param *CategoryDO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuCategoryQueryAPIRequest) GetParam() *CategoryDO {
    return r._param
}
