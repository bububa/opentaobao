package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建买赠活动 APIRequest
alibaba.wdk.marketing.itembuygift.createactivity

创建买赠活动
*/
type AlibabaWdkMarketingItembuygiftCreateactivityRequest struct {
    model.Params

    // 创建活动请求入参
    param   *ItemBuyGiftActivity 

}

func NewAlibabaWdkMarketingItembuygiftCreateactivityRequest() *AlibabaWdkMarketingItembuygiftCreateactivityRequest{
    return &AlibabaWdkMarketingItembuygiftCreateactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItembuygiftCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itembuygift.createactivity"
}

func (r AlibabaWdkMarketingItembuygiftCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItembuygiftCreateactivityRequest) SetParam(param *ItemBuyGiftActivity) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItembuygiftCreateactivityRequest) GetParam() *ItemBuyGiftActivity {
    return r.param
}

