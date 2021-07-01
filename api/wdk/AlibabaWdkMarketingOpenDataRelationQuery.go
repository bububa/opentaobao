package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
数据关联关系查询 
alibaba.wdk.marketing.open.data.relation.query

数据关联关系查询
*/
func AlibabaWdkMarketingOpenDataRelationQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenDataRelationQueryAPIRequest, session string) (*wdk.AlibabaWdkMarketingOpenDataRelationQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingOpenDataRelationQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
