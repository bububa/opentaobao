package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品 API请求
alibaba.wdk.sku.query

查询商品
*/
type AlibabaWdkSkuQueryRequest struct {
    model.Params
    // 入参
    _param   *SkuQueryDo
}

// 初始化AlibabaWdkSkuQueryRequest对象
func NewAlibabaWdkSkuQueryRequest() *AlibabaWdkSkuQueryRequest{
    return &AlibabaWdkSkuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaWdkSkuQueryRequest) SetParam(_param *SkuQueryDo) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuQueryRequest) GetParam() *SkuQueryDo {
    return r._param
}
