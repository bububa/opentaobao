package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterWorkcardCreateAPIRequest
服务平台工单创建接口 API请求
alibaba.servicecenter.workcard.create

创建服务平台工单 */
type AlibabaServicecenterWorkcardCreateAPIRequest struct {
	model.Params
	// 服务单id
	_spServiceOrderId int64
	// 申请工单时的序号，对应服务单上的serviceSequence。用于控制幂等，防重复提交
	_serviceSequence int64
	// 申请次数
	_serviceCount int64
	// 工单属性，json格式字符串
	_attributes string
	// 工单外部唯一键单号
	_outerId string
	// 服务提供者信息
	_serviceProvider *ServiceProviderDto
}

// New
