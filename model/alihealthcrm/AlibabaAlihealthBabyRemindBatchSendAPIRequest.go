package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBabyRemindBatchSendAPIRequest
批量发送母婴提醒 API请求
alibaba.alihealth.baby.remind.batch.send

批量发送母婴提醒 */
type AlibabaAlihealthBabyRemindBatchSendAPIRequest struct {
	model.Params
	// 批量发送提醒
	_batchRemindRequests []BatchRemindRequestDto
	// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
	_remindType int64
}

// New
