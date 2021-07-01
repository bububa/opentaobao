package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest
ISV获取报告文件查看验证码 API请求
alibaba.alihealth.examination.report.diagnose.file.code.get

体检报告人工解读_ISV获取报告文件验证码进行查看报告文件 */
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest struct {
	model.Params
	// 报告id
	_reportId int64
	// 订单id
	_orderId string
	// 医生id
	_doctorId string
}

// New
