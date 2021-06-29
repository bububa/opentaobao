package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商品 API请求
alibaba.wdk.sku.channelsku.query

查询渠道商品
*/
type AlibabaWdkSkuChannelskuQueryRequest struct {
    model.Params
    // 查询渠道商品的入参
    _param   *ChannelSkuQueryDo
}

// 初始化AlibabaWdkSkuChannelskuQueryRequest对象
func NewAlibabaWdkSkuChannelskuQueryRequest() *AlibabaWdkSkuChannelskuQueryRequest{
    return &AlibabaWdkSkuChannelskuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询渠道商品的入参
func (r *AlibabaWdkSkuChannelskuQueryRequest) SetParam(_param *ChannelSkuQueryDo) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuChannelskuQueryRequest) GetParam() *ChannelSkuQueryDo {
    return r._param
}
