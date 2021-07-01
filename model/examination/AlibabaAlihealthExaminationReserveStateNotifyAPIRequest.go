package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveStateNotifyAPIRequest
体检机构对接_体检状态主动通知 API请求
alibaba.alihealth.examination.reserve.state.notify

到了体检当天后，服务商主动通知体检预约状态 */
type AlibabaAlihealthExaminationReserveStateNotifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 健康预约凭证
	_reserveNumber string
	// 体检状态：未到检(exam_not), 已到检(exam_done)； 上门服务中还需以下两种状态：预约确认中（reserve_confirming），预约拒绝（reserve_rejected）；
	_reportStatus string
	// 到检凭证，exam_done状态下，该字段必填
	_checkNo string
}

// New
