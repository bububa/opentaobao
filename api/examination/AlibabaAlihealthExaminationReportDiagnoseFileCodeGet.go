package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
ISV获取报告文件查看验证码 
alibaba.alihealth.examination.report.diagnose.file.code.get

体检报告人工解读_ISV获取报告文件验证码进行查看报告文件
*/
func AlibabaAlihealthExaminationReportDiagnoseFileCodeGet(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest, session string) (*examination.AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
