package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceCreateResultGetAPIRequest
ERP开票结果获取 API请求
alibaba.einvoice.create.result.get

ERP开票结果获取 */
type AlibabaEinvoiceCreateResultGetAPIRequest struct {
	model.Params
	// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
	_serialNo string
	// 电商平台代码。淘宝：taobao，天猫：tmall
	_platformCode string
	// 电商平台对应的订单号
	_platformTid string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
	_outShopName string
}

// New
