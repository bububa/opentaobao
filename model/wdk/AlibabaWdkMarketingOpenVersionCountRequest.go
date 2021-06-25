package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
版本数量查询 APIRequest
alibaba.wdk.marketing.open.version.count

版本数量查询
*/
type AlibabaWdkMarketingOpenVersionCountRequest struct {
    model.Params

    // 查询版本号
    versionId   int64 

    // 操作Id
    operateId   string 

}

func NewAlibabaWdkMarketingOpenVersionCountRequest() *AlibabaWdkMarketingOpenVersionCountRequest{
    return &AlibabaWdkMarketingOpenVersionCountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingOpenVersionCountRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.version.count"
}

func (r AlibabaWdkMarketingOpenVersionCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingOpenVersionCountRequest) SetVersionId(versionId int64) error {
    r.versionId = versionId
    r.Set("version_id", versionId)
    return nil
}

func (r AlibabaWdkMarketingOpenVersionCountRequest) GetVersionId() int64 {
    return r.versionId
}

func (r *AlibabaWdkMarketingOpenVersionCountRequest) SetOperateId(operateId string) error {
    r.operateId = operateId
    r.Set("operate_id", operateId)
    return nil
}

func (r AlibabaWdkMarketingOpenVersionCountRequest) GetOperateId() string {
    return r.operateId
}

