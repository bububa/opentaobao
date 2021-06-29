package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交版本号 API请求
alibaba.wdk.marketing.version.commit

提交版本号，标识结束此版本操作
*/
type AlibabaWdkMarketingVersionCommitRequest struct {
    model.Params
    // 版本号提交参数
    param   *SeasonVersionCommitParam
}

// 初始化AlibabaWdkMarketingVersionCommitRequest对象
func NewAlibabaWdkMarketingVersionCommitRequest() *AlibabaWdkMarketingVersionCommitRequest{
    return &AlibabaWdkMarketingVersionCommitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingVersionCommitRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.version.commit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingVersionCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 版本号提交参数
func (r *AlibabaWdkMarketingVersionCommitRequest) SetParam(param *SeasonVersionCommitParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingVersionCommitRequest) GetParam() *SeasonVersionCommitParam {
    return r.param
}
