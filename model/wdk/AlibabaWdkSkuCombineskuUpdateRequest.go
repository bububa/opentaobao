package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品更新接口 API请求
alibaba.wdk.sku.combinesku.update

组合商品修改接口
*/
type AlibabaWdkSkuCombineskuUpdateAPIRequest struct {
    model.Params
    // 请求参数
    _paramList   []SkuDO
}

// 初始化AlibabaWdkSkuCombineskuUpdateAPIRequest对象
func NewAlibabaWdkSkuCombineskuUpdateRequest() *AlibabaWdkSkuCombineskuUpdateAPIRequest{
    return &AlibabaWdkSkuCombineskuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuUpdateAPIRequest) SetParamList(_paramList []SkuDO) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetParamList() []SkuDO {
    return r._paramList
}
