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
type AlibabaWdkSkuCombineskuAddRequest struct {
    model.Params
    // 请求参数
    paramList   []SkuDo
}

// 初始化AlibabaWdkSkuCombineskuAddRequest对象
func NewAlibabaWdkSkuCombineskuAddRequest() *AlibabaWdkSkuCombineskuAddRequest{
    return &AlibabaWdkSkuCombineskuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuAddRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuCombineskuAddRequest) GetParamList() []SkuDo {
    return r.paramList
}
