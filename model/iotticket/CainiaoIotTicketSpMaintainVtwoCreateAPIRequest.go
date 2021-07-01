package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpMaintainVtwoCreateAPIRequest
服务商制定维修费方案 API请求
cainiao.iot.ticket.sp.maintain.vtwo.create

服务商制定维修费方案 */
type CainiaoIotTicketSpMaintainVtwoCreateAPIRequest struct {
	model.Params
	// 维修方案
	_makeMaintainPlanTopRequest *MakeMaintainPlanV2TopRequest
}

// New
