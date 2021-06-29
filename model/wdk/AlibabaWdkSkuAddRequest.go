package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增商品 API请求
alibaba.wdk.sku.add

创建RT门店商品或DC商品
*/
type AlibabaWdkSkuAddRequest struct {
    model.Params
    // 商品列表
    paramList   []SkuDo
}

// 初始化AlibabaWdkSkuAddRequest对象
func NewAlibabaWdkSkuAddRequest() *AlibabaWdkSkuAddRequest{
    return &AlibabaWdkSkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 商品列表
func (r *AlibabaWdkSkuAddRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuAddRequest) GetParamList() []SkuDo {
    return r.paramList
}
