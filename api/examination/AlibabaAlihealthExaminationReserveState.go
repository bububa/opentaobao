package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
体检机构对接_体检状态查询 
alibaba.alihealth.examination.reserve.state

体检机构对接_体检状态查询
*/
func AlibabaAlihealthExaminationReserveState(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveStateRequest, session string) (*examination.AlibabaAlihealthExaminationReserveStateAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReserveStateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
