package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
体检机构对接_预约取消 
alibaba.alihealth.examination.reserve.cancel

体检机构对接_体检取消
*/
func AlibabaAlihealthExaminationReserveCancel(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveCancelAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReserveCancelAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReserveCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
