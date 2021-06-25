package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建活动新接口 APIRequest
alibaba.wdk.marketing.itempool.activity.create

创建活动新接口，支持新工具玩法
*/
type AlibabaWdkMarketingItempoolActivityCreateRequest struct {
    model.Params

    // 创建活动请求入参
    param   *ItemPoolActivity 

}

func NewAlibabaWdkMarketingItempoolActivityCreateRequest() *AlibabaWdkMarketingItempoolActivityCreateRequest{
    return &AlibabaWdkMarketingItempoolActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingItempoolActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.activity.create"
}

func (r AlibabaWdkMarketingItempoolActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingItempoolActivityCreateRequest) SetParam(param *ItemPoolActivity) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingItempoolActivityCreateRequest) GetParam() *ItemPoolActivity {
    return r.param
}

