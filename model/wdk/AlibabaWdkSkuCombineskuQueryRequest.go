package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品查询接口 API请求
alibaba.wdk.sku.combinesku.query

查询组合商品接口
*/
type AlibabaWdkSkuCombineskuQueryRequest struct {
    model.Params
    // 请求参数
    _param   *SkuQueryDo
}

// 初始化AlibabaWdkSkuCombineskuQueryRequest对象
func NewAlibabaWdkSkuCombineskuQueryRequest() *AlibabaWdkSkuCombineskuQueryRequest{
    return &AlibabaWdkSkuCombineskuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuQueryRequest) SetParam(_param *SkuQueryDo) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuCombineskuQueryRequest) GetParam() *SkuQueryDo {
    return r._param
}
