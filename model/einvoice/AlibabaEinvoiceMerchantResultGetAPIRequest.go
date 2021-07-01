package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceMerchantResultGetAPIRequest
商家自研ERP开票结果获取 API请求
alibaba.einvoice.merchant.result.get

商家自研ERP开票结果获取 */
type AlibabaEinvoiceMerchantResultGetAPIRequest struct {
	model.Params
	// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
	_serialNo string
	// 电商平台代码。淘宝：taobao，天猫：tmall
	_platformCode string
	// 电商平台对应的订单号
	_platformTid string
	// 收款方税务登记证号
	_payeeRegisterNo string
}

// New
