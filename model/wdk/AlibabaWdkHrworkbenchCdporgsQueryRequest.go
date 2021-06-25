package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
homs人力资源组织树查询接口 APIRequest
alibaba.wdk.hrworkbench.cdporgs.query

提供查询homs人力组织树的接口，按照商家做权限隔离。
*/
type AlibabaWdkHrworkbenchCdporgsQueryRequest struct {
    model.Params

}

func NewAlibabaWdkHrworkbenchCdporgsQueryRequest() *AlibabaWdkHrworkbenchCdporgsQueryRequest{
    return &AlibabaWdkHrworkbenchCdporgsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkHrworkbenchCdporgsQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.hrworkbench.cdporgs.query"
}

func (r AlibabaWdkHrworkbenchCdporgsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


