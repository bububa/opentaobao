package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveIsvModifyAPIRequest
ISV调TOP主动发起改期信息 API请求
alibaba.alihealth.examination.reserve.isv.modify

体检机构对接_ISV发起体检套餐改期 */
type AlibabaAlihealthExaminationReserveIsvModifyAPIRequest struct {
	model.Params
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
	// 修改后预约预约日期，格式：yyyy-MM-dd
	_reserveDate string
	// 修改后预约时间段开始时间 HH:mm:ss
	_reserveTimeStart string
	// 修改后预约时间段结束时间 HH:mm:ss
	_reserveTimeEnd string
}

// New
