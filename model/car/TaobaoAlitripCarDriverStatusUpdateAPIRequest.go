package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarDriverStatusUpdateAPIRequest
司机服务状态更新接口 API请求
taobao.alitrip.car.driver.status.update

飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态 */
type TaobaoAlitripCarDriverStatusUpdateAPIRequest struct {
	model.Params
	// 飞猪订单id
	_orderId string
	// 服务商订单id
	_thirdOrderId string
	// 服务商标识，由飞猪提供给到各服务商
	_providerId string
	// 司机服务状态。1-司机已出发，2-司机已到达，3-司机已开始服务
	_status int64
	// 状态变更相应时间（如司机出发时间、司机到达时间、服务开始时间），格式：yyyy-mm-dd hh:mm:ss
	_time string
	// 可选，卖家id
	_sellerId string
	// 0:接送机 1：实时打车 2：租车(不传值默认为0)
	_useType int64
}

// New
