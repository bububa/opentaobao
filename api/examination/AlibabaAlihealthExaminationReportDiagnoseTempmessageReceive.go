package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
导医通报告解读临时消息接收 
alibaba.alihealth.examination.report.diagnose.tempmessage.receive

导医通报告解读临时消息接收
*/
func AlibabaAlihealthExaminationReportDiagnoseTempmessageReceive(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest, session string) (*examination.AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
