package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest
上门检测服务信息同步 API请求
alibaba.alihealth.examination.todoor.serviceinfo.sync

isv同步上门检测服务信息给健康 */
type AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 从业者信息
	_medicalPractitionerInfo *MedicalPractitionerInfo
	// 健康预约凭证
	_reserveNumber string
	// 事件(ASSIGNED_PRACTITONER:已分配医护人员、PRACTITONER_GO_OUT:医护人员已出发、PRACTITONER_HOME:医护人员已到家、PRACTITONER_CHECKED:医护人员检查完成)、CHANGE_PRACTITONER(变更医护人员)
	_event string
	// 事件发生时间
	_eventOccurTime string
}

// New
