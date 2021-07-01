package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePartnerReturnAPIRequest
开票商回传开票结果 API请求
alibaba.einvoice.partner.return

开票商返回开票结果数据 */
type AlibabaEinvoicePartnerReturnAPIRequest struct {
	model.Params
	// 电商平台身份标识码，TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
	_platformCode string
	// 开票金额
	_invoiceAmount string
	// 发票密文，密码区的字符串
	_ciphertext string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 二维码
	_qrCode string
	// erp自定义单据号
	_erpTid string
	// 文件类型(pdf,jpg,png)
	_fileDataType string
	// 发票号码
	_invoiceNo string
	// 发票日期
	_invoiceDate string
	// 发票文件PDF内容，PDF的byte[]字段串。
	_invoiceFileData *model.File
	// 流水号
	_serialNo string
	// 防伪码
	_antiFakeCode string
	// 税控设备编号(新版电子发票有)
	_deviceNo string
	// 发票代码
	_invoiceCode string
	// 电商平台对应的订单号
	_platformTid string
	// 开票结果"success"或者"fail"
	_createResult string
	// 错误码
	_bizErrorCode string
	// 错误信息
	_bizErrorMsg string
	// 开票请求的唯一索引
	_reqIndex string
	// 开票时间，格式为HH:mm:ss
	_invoiceTime string
}

// New
