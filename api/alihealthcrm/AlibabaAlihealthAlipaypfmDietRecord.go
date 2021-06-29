package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
用户每日摄入卡路里总量回传接口 
alibaba.alihealth.alipaypfm.diet.record

用户每日摄入卡路里总量回传接口
*/
func AlibabaAlihealthAlipaypfmDietRecord(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmDietRecordRequest, session string) (*alihealthcrm.AlibabaAlihealthAlipaypfmDietRecordAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthAlipaypfmDietRecordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
