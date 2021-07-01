package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceBillEinvoiceListAPIRequest
扫码开票列表 API请求
alibaba.einvoice.bill.einvoice.list

扫码开票列表，包括用户扫二维码开票和结算单同步前的开票数据 */
type AlibabaEinvoiceBillEinvoiceListAPIRequest struct {
	model.Params
	// 结算单同步的ERP平台系统
	_platform string
	// 收款方税号
	_payeeRegisterNo string
	// 订单ID
	_orderId string
	// 开票状态：0=未开票，1=开票中，3=开蓝成功，4=开蓝失败。不填获取全部
	_einvoiceType []int64
}

// New
