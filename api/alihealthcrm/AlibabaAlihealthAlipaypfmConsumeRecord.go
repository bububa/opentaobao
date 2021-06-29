package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
记录用户每日消耗卡路里总量 
alibaba.alihealth.alipaypfm.consume.record

记录用户每日消耗卡路里总量
*/
func AlibabaAlihealthAlipaypfmConsumeRecord(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmConsumeRecordRequest, session string) (*alihealthcrm.AlibabaAlihealthAlipaypfmConsumeRecordAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthAlipaypfmConsumeRecordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
