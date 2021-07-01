package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceMerchantBindcompanyAPIRequest
发票中台-跨平台绑定已入驻税号与商户 API请求
alibaba.einvoice.merchant.bindcompany

税号在阿里发票平台入驻成功后，允许业务方通过本接口跨业务平台绑定入驻税号和业务平台商户，绑定成功后该商户可以使用该税号的盘进行开票。绑定成功后，可以使用同平台授权、取消授权税号适用商户接口来变更税号和商户关系。 */
type AlibabaEinvoiceMerchantBindcompanyAPIRequest struct {
	model.Params
	// 业务方发起首次绑定门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
	_outerId string
	// 业务平台商户ID
	_merchantUserId string
	// 激活码
	_activationCode string
	// 业务平台code, 由阿里发票分配
	_platformCode string
	// 税务登记号
	_payeeRegisterNo string
	// 业务平台门店名称
	_merchantName string
	// 税号已入驻的原业务平台code
	_sourcePlatformCode string
}

// New
