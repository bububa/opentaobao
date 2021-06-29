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
type AlibabaWdkSkuChannelskuUpdateRequest struct {
    model.Params
    // 请求参数
    _paramList   []ChannelSkuDo
}

// 初始化AlibabaWdkSkuChannelskuUpdateRequest对象
func NewAlibabaWdkSkuChannelskuUpdateRequest() *AlibabaWdkSkuChannelskuUpdateRequest{
    return &AlibabaWdkSkuChannelskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuChannelskuUpdateRequest) SetParamList(_paramList []ChannelSkuDo) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuChannelskuUpdateRequest) GetParamList() []ChannelSkuDo {
    return r._paramList
}
