package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
服务商主动通知体检报告 
alibaba.alihealth.examination.reserve.report.nofify

服务商主动回传用户的体检报告数据
*/
func AlibabaAlihealthExaminationReserveReportNofify(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveReportNofifyAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReserveReportNofifyAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReserveReportNofifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
