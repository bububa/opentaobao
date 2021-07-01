package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
体检机构对接_体检状态主动通知 
alibaba.alihealth.examination.reserve.state.notify

到了体检当天后，服务商主动通知体检预约状态
*/
func AlibabaAlihealthExaminationReserveStateNotify(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveStateNotifyAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReserveStateNotifyAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReserveStateNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
