package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新渠道商品 API请求
alibaba.wdk.sku.channelsku.update

批量更新渠道商品，商家通过Top接入
*/
type AlibabaWdkSkuChannelskuUpdateAPIRequest struct {
    model.Params
    // 请求参数
    _paramList   []ChannelSkuDO
}

// 初始化AlibabaWdkSkuChannelskuUpdateAPIRequest对象
func NewAlibabaWdkSkuChannelskuUpdateRequest() *AlibabaWdkSkuChannelskuUpdateAPIRequest{
    return &AlibabaWdkSkuChannelskuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuChannelskuUpdateAPIRequest) SetParamList(_paramList []ChannelSkuDO) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetParamList() []ChannelSkuDO {
    return r._paramList
}
