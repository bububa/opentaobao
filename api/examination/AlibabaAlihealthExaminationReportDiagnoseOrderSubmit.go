package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
体检报告人工解读订单 
alibaba.alihealth.examination.report.diagnose.order.submit

体检报告人工解读订单信息推送给ISV，进行人工解读
*/
func AlibabaAlihealthExaminationReportDiagnoseOrderSubmit(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest, session string) (*examination.AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
