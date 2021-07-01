package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceBillSyncAPIRequest
结算单同步 API请求
alibaba.einvoice.bill.sync

电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票 */
type AlibabaEinvoiceBillSyncAPIRequest struct {
	model.Params
	// 结算商品单明细列表
	_invoiceItems []BillItemDo
	// 结算单同步操作：=1插入，=2更新，=3废弃删除
	_status int64
	// 结算单订单日期
	_orderDate string
	// 店铺名称，与后台店铺名称保持一致
	_shopName string
	// 税务登记证号
	_payeeRegisterNo string
	// 结算单订单ID
	_orderId string
	// 结算单总价格，小数点后2两位
	_sumPrice string
	// 调用平台，用于区分同一个税号下多个店铺来源["TB:淘宝","ALIPAY:支付宝","TM:天猫","JD:京东","DD:当当","PP:拍拍","YX:易讯","EBAY:ebay","QQ:QQ网购","AMAZON:亚马逊","SN:苏宁","GM:国美","WPH:唯品会","JM:聚美","LF:乐蜂","MGJ:蘑菇街","JS:聚尚","PX:拍鞋","YT:银泰","YHD:1号店","VANCL:凡客","YL:邮乐","YG:优购","1688:阿里巴巴","POS:POS门店","ELEME:饿了么","OTHER:其他"]
	_platform string
	// 生成二维码参数，若不需要生成二维码，则不填
	_qrcode *QrCodeDo
	// 品牌名称，不填默认=shop_name
	_brandName string
	// 结算单可开票总金额（不填=sumPrice），小数点后2两位
	_invoicePrice string
	// 开票店铺的平台，默认等于platform
	_shopPlatform string
}

// New
