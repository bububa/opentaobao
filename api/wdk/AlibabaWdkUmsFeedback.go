package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
质量反馈（入库辅助）-ERP下发单 
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单
*/
func AlibabaWdkUmsFeedback(clt *core.SDKClient, req *wdk.AlibabaWdkUmsFeedbackAPIRequest, session string) (*wdk.AlibabaWdkUmsFeedbackAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
