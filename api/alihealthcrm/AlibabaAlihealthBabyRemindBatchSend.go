package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
批量发送母婴提醒 
alibaba.alihealth.baby.remind.batch.send

批量发送母婴提醒
*/
func AlibabaAlihealthBabyRemindBatchSend(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthBabyRemindBatchSendRequest, session string) (*alihealthcrm.AlibabaAlihealthBabyRemindBatchSendAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthBabyRemindBatchSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}