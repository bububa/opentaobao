package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest
阿里健康医保支付信息获取 API请求
alibaba.alihealth.nr.trade.medical.insurance.get

阿里健康医保支付信息获取 */
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// New
