package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest
上传订单同城快递单号 API请求
alibaba.health.nr.logistics.deliveryno.update

上传订单同城快递单号 */
type AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 快递公司代码
	_cpCode string
	// 快递单号
	_courierNo string
	// 是否强制上传，1代表强制，其他值代表需要进行cp_code合法性校验
	_force int64
}

// New
