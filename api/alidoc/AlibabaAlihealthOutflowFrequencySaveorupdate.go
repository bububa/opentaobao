package alidoc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alidoc"
)

/* 
处方外流-药品频次同步接口 
alibaba.alihealth.outflow.frequency.saveorupdate

处方外流-药品频次同步接口
*/
func AlibabaAlihealthOutflowFrequencySaveorupdate(clt *core.SDKClient, req *alidoc.AlibabaAlihealthOutflowFrequencySaveorupdateRequest, session string) (*alidoc.AlibabaAlihealthOutflowFrequencySaveorupdateAPIResponse, error) {
    var resp alidoc.AlibabaAlihealthOutflowFrequencySaveorupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
