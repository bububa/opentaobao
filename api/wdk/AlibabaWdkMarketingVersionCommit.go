package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
提交版本号 
alibaba.wdk.marketing.version.commit

提交版本号，标识结束此版本操作
*/
func AlibabaWdkMarketingVersionCommit(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingVersionCommitRequest, session string) (*wdk.AlibabaWdkMarketingVersionCommitResponse, error) {
    var resp wdk.AlibabaWdkMarketingVersionCommitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
