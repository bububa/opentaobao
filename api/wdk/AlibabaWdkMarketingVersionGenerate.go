package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
生成发布使用的版本号 
alibaba.wdk.marketing.version.generate

生成发布使用的版本号
*/
func AlibabaWdkMarketingVersionGenerate(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingVersionGenerateRequest, session string) (*wdk.AlibabaWdkMarketingVersionGenerateAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingVersionGenerateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
