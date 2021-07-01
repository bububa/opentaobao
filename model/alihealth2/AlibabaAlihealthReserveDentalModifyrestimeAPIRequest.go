package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalModifyrestimeAPIRequest
修改预约时间 API请求
alibaba.alihealth.reserve.dental.modifyrestime

修改预约时间 */
type AlibabaAlihealthReserveDentalModifyrestimeAPIRequest struct {
	model.Params
	// 预约单ID
	_reserveId int64
	// 预约时间
	_reserveTime string
}

// New
