package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方外流-问诊、处方同步 
alibaba.alihealth.outflow.visitinfo.sync

阿里健康-处方外流-对外提供问诊、处方功能
*/
func AlibabaAlihealthOutflowVisitinfoSync(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowVisitinfoSyncAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowVisitinfoSyncAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowVisitinfoSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
