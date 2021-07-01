package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceQrcodeCreateAPIRequest
扫码开票二维码生成 API请求
alibaba.einvoice.qrcode.create

扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码 */
type AlibabaEinvoiceQrcodeCreateAPIRequest struct {
	model.Params
	// 发票商品明细
	_invoiceItems []BillItemDo
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 开票的订单号，同结算单订单号
	_orderId string
	// 开票金额
	_sumPrice string
	// 请求方唯一标识ID，例如POS机编码
	_sourceId string
	// 二维码图片中间的logo
	_qrLogo string
	// 二维码图片宽度，默认=450
	_width int64
	// 二维码图片高度，默认=450
	_height int64
	// 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
	_qrType int64
	// 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
	_platform string
}

// New
