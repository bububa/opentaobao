package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
ISV调TOP主动发起改期信息 
alibaba.alihealth.examination.reserve.isv.modify

体检机构对接_ISV发起体检套餐改期
*/
func AlibabaAlihealthExaminationReserveIsvModify(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveIsvModifyAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReserveIsvModifyAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReserveIsvModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
