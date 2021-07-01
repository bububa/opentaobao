package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest
导医通报告解读临时消息接收 API请求
alibaba.alihealth.examination.report.diagnose.tempmessage.receive

导医通报告解读临时消息接收 */
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest struct {
	model.Params
	// 入参对象
	_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest
}

// New
