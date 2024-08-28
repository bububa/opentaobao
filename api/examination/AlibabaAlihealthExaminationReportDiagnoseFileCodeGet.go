package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseFileCodeGet ISV获取报告文件查看验证码
// alibaba.alihealth.examination.report.diagnose.file.code.get
//
// 体检报告人工解读_ISV获取报告文件验证码进行查看报告文件
func AlibabaAlihealthExaminationReportDiagnoseFileCodeGet(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
