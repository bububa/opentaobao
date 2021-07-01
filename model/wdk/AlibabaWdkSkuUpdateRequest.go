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
type AlibabaWdkSkuUpdateAPIRequest struct {
    model.Params
    // 参数列表
    _paramList   []SkuDO
}

// 初始化AlibabaWdkSkuUpdateAPIRequest对象
func NewAlibabaWdkSkuUpdateRequest() *AlibabaWdkSkuUpdateAPIRequest{
    return &AlibabaWdkSkuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 参数列表
func (r *AlibabaWdkSkuUpdateAPIRequest) SetParamList(_paramList []SkuDO) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuUpdateAPIRequest) GetParamList() []SkuDO {
    return r._paramList
}
