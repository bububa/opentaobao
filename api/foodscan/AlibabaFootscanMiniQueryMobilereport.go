package foodscan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/foodscan"
)

/* 
根据scanId查询报告 
alibaba.footscan.mini.query.mobilereport

根据scanId查询报告
*/
func AlibabaFootscanMiniQueryMobilereport(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniQueryMobilereportRequest, session string) (*foodscan.AlibabaFootscanMiniQueryMobilereportAPIResponse, error) {
    var resp foodscan.AlibabaFootscanMiniQueryMobilereportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}