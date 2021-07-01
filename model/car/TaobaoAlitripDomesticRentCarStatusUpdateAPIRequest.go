package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest
航旅国内租车订单状态更新 API请求
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新 */
type TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest struct {
	model.Params
	// 121-用车中（用户取车成功） 122-待结算（用户还车成功）
	_status int64
	// 服务商平台订单号
	_thirdOrderId string
	// 飞猪平台订单号
	_orderId string
	// 服务商标识，由飞猪提供给到各服务商
	_providerId string
}

// New
