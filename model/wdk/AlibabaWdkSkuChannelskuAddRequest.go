package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增渠道商品 API请求
alibaba.wdk.sku.channelsku.add

盒马帮1期新增渠道商品
*/
type AlibabaWdkSkuChannelskuAddRequest struct {
    model.Params
    // 入参模型
    chSkuDOList   []ChannelSkuDo
}

// 初始化AlibabaWdkSkuChannelskuAddRequest对象
func NewAlibabaWdkSkuChannelskuAddRequest() *AlibabaWdkSkuChannelskuAddRequest{
    return &AlibabaWdkSkuChannelskuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.channelsku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChSkuDOList Setter
// 入参模型
func (r *AlibabaWdkSkuChannelskuAddRequest) SetChSkuDOList(chSkuDOList []ChannelSkuDo) error {
    r.chSkuDOList = chSkuDOList
    r.Set("ch_sku_d_o_list", chSkuDOList)
    return nil
}

// ChSkuDOList Getter
func (r AlibabaWdkSkuChannelskuAddRequest) GetChSkuDOList() []ChannelSkuDo {
    return r.chSkuDOList
}
