package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest
体检报告人工解读订单 API请求
alibaba.alihealth.examination.report.diagnose.order.submit

体检报告人工解读订单信息推送给ISV，进行人工解读 */
type AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest struct {
	model.Params
	// 订单ID
	_orderId string
	// 手机号码，显示后四位
	_mobilePhone string
	// 证件号，显示前1，后1
	_idCardNo string
	// 性别
	_gender string
	// 报告地址
	_reportUrl string
	// 主诉问题
	_question string
	// 咨询人名称
	_patientName string
}

// New
