package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品池活动 APIRequest
alibaba.wdk.marketing.itempool.createactivity

添加商品池活动
*/
type AlibabaWdkMarketingItempoolCreateactivityRequest struct {
    model.Params

    // 创建活动请求入参
    param   *ItemPoolActivity 

}

func NewAlibabaWdkMarketingItempoolCreateactivityRequest() *AlibabaWdkMarketingItempoolCreateactivityRequest{
    return &AlibabaWdkMarketingItempoolCreateactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.createactivity"
}

func (r AlibabaWdkMarketingItempoolCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolCreateactivityRequest) SetParam(param *ItemPoolActivity) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItempoolCreateactivityRequest) GetParam() *ItemPoolActivity {
    return r.param
}

