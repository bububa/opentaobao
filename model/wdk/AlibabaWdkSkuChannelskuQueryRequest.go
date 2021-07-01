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
type AlibabaWdkSkuChannelskuQueryAPIRequest struct {
    model.Params
    // 查询渠道商品的入参
    _param   *ChannelSkuQueryDO
}

// 初始化AlibabaWdkSkuChannelskuQueryAPIRequest对象
func NewAlibabaWdkSkuChannelskuQueryRequest() *AlibabaWdkSkuChannelskuQueryAPIRequest{
    return &AlibabaWdkSkuChannelskuQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 查询渠道商品的入参
func (r *AlibabaWdkSkuChannelskuQueryAPIRequest) SetParam(_param *ChannelSkuQueryDO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetParam() *ChannelSkuQueryDO {
    return r._param
}
