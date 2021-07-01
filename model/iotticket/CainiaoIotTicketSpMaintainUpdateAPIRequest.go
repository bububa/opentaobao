package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpMaintainUpdateAPIRequest
IoT售后服务商维修方案更新 API请求
cainiao.iot.ticket.sp.maintain.update

IoT售后服务商维修方案更新 */
type CainiaoIotTicketSpMaintainUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *UpdateMaintainPlanTopRequest
}

// New
