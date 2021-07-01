package alihealthcert

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest
健康证服务商预约结果通知阿里健康 API请求
alibaba.alihealth.examination.reserve.certificate.notify

当ISV执行完健康证预约成功之后， 调用通知阿里健康 */
type AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 健康预约凭证
	_reserveNumber string
	// 预约成功（ISSUED），预约失败(FAIL)
	_reportStatus string
	// 到检凭证，exam_done状态下，该字段必填
	_checkNo string
	// 预约状态原因，描述预约状态如FAIL时的原因
	_statusReason string
}

// New
