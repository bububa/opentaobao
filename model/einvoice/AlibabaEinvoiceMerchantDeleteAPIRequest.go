package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceMerchantDeleteAPIRequest
发票中台-同平台取消授权税号适用商户 API请求
alibaba.einvoice.merchant.delete

税号授权给同平台下其他商户使用后，可以使用此接口取消授权，被取消授权的商户失去开票能力 */
type AlibabaEinvoiceMerchantDeleteAPIRequest struct {
	model.Params
	// 验证码，商户首次绑定已入驻税号接口返回的taxToken
	_taxToken string
	// 业务方发起删除商户的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
	_outerId string
	// 业务平台商户ID
	_merchantUserId string
	// 业务平台code, 由阿里发票分配
	_platformCode string
	// 税号
	_payeeRegisterNo string
}

// New
