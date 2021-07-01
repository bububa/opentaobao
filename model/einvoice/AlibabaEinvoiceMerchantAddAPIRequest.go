package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceMerchantAddAPIRequest
发票中台-同平台授权税号适用商户 API请求
alibaba.einvoice.merchant.add

适用于以下场景：
业务税号入驻成功后，需要将税号授权给同平台下其他商户，使得其他商户也具备开票能力 */
type AlibabaEinvoiceMerchantAddAPIRequest struct {
	model.Params
	// 验证码，门店绑定已入驻税号接口返回的taxToken
	_taxToken string
	// 业务方发起新增门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
	_outerId string
	// 业务平台门店ID
	_merchantUserId string
	// 业务平台code, 由阿里发票分配
	_platformCode string
	// 税务登记号
	_payeeRegisterNo string
	// 业务平台门店名称
	_merchantName string
	// 税盘列表
	_deviceIds []string
}

// New
