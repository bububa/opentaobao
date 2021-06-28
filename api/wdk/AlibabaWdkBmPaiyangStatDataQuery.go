package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
派样统计数据查询 
alibaba.wdk.bm.paiyang.stat.data.query

派样统计数据查询
*/
func AlibabaWdkBmPaiyangStatDataQuery(clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangStatDataQueryRequest, session string) (*wdk.AlibabaWdkBmPaiyangStatDataQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkBmPaiyangStatDataQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
