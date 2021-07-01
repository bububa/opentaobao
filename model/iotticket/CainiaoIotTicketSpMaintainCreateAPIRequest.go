package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpMaintainCreateAPIRequest
IoT售后服务商制定维修方案 API请求
cainiao.iot.ticket.sp.maintain.create

IoT售后服务商制定维修方案 */
type CainiaoIotTicketSpMaintainCreateAPIRequest struct {
	model.Params
	// 请求参数
	_param *AssignMaintenancePersonnelTopRequest
}

// New
