package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicemonitormessageUpdateAPIRequest
服务商更新预警消息状态 API请求
tmall.servicecenter.servicemonitormessage.update

服务商收到预警后，需要进行回复已读状态，并可填写备注 */
type TmallServicecenterServicemonitormessageUpdateAPIRequest struct {
	model.Params
	// 预警消息id
	_serviceMonitorMessageId int64
	// 预警处理备注
	_memo string
	// 可更新状态：3、已读
	_status int64
}

// New
