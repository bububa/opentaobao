package foodscan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/foodscan"
)

/* 
第二只脚生成报告接口 
alibaba.footscan.mini.report.fragment.second

第二只脚生成报告接口
*/
func AlibabaFootscanMiniReportFragmentSecond(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniReportFragmentSecondRequest, session string) (*foodscan.AlibabaFootscanMiniReportFragmentSecondAPIResponse, error) {
    var resp foodscan.AlibabaFootscanMiniReportFragmentSecondAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
