package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品 API请求
alibaba.wdk.sku.update

开放商品更新接口
*/
type AlibabaWdkSkuUpdateRequest struct {
    model.Params
    // 参数列表
    _paramList   []SkuDO
}

// 初始化AlibabaWdkSkuUpdateRequest对象
func NewAlibabaWdkSkuUpdateRequest() *AlibabaWdkSkuUpdateRequest{
    return &AlibabaWdkSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 参数列表
func (r *AlibabaWdkSkuUpdateRequest) SetParamList(_paramList []SkuDO) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuUpdateRequest) GetParamList() []SkuDO {
    return r._paramList
}
