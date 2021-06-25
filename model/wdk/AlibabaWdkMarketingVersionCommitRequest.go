package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交版本号 APIRequest
alibaba.wdk.marketing.version.commit

提交版本号，标识结束此版本操作
*/
type AlibabaWdkMarketingVersionCommitRequest struct {
    model.Params

    // 版本号提交参数
    param   *SeasonVersionCommitParam 

}

func NewAlibabaWdkMarketingVersionCommitRequest() *AlibabaWdkMarketingVersionCommitRequest{
    return &AlibabaWdkMarketingVersionCommitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingVersionCommitRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.version.commit"
}

func (r AlibabaWdkMarketingVersionCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingVersionCommitRequest) SetParam(param *SeasonVersionCommitParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingVersionCommitRequest) GetParam() *SeasonVersionCommitParam {
    return r.param
}

