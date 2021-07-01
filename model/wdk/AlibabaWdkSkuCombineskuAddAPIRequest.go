package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品新增接口 API请求
alibaba.wdk.sku.combinesku.add

组合商品新增接口
*/
type AlibabaWdkSkuCombineskuAddAPIRequest struct {
    model.Params
    // 请求参数
    _paramList   []SkuDo
}

// 初始化AlibabaWdkSkuCombineskuAddAPIRequest对象
func NewAlibabaWdkSkuCombineskuAddRequest() *AlibabaWdkSkuCombineskuAddAPIRequest{
    return &AlibabaWdkSkuCombineskuAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuAddAPIRequest) SetParamList(_paramList []SkuDo) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetParamList() []SkuDo {
    return r._paramList
}
