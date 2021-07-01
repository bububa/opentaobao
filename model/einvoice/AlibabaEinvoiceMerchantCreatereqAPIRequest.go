package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceMerchantCreatereqAPIRequest
商家自研ERP开票请求接口 API请求
alibaba.einvoice.merchant.createreq

商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票 */
type AlibabaEinvoiceMerchantCreatereqAPIRequest struct {
	model.Params
	// 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
	_businessType int64
	// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
	_erpTid string
	// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
	_platformCode string
	// 电商平台对应的主订单号
	_platformTid string
	// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
	_serialNo string
	// 开票方地址(新版中为必传)
	_payeeAddress string
	// 开票方银行及 帐号
	_payeeBankaccount string
	// 开票方名称，公司名(如:XX商城)
	_payeeName string
	// 付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
	_payerRegisterNo string
	// 开票人
	_payeeOperator string
	// 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
	_invoiceAmount string
	// 电子发票明细
	_invoiceItems []InvoiceItem
	// 发票备注，有些省市会把此信息打印到PDF中
	_invoiceMemo string
	// 发票(开票)类型，蓝票blue,红票red，默认blue
	_invoiceType string
	// 原发票代码(开红票时传入)
	_normalInvoiceCode string
	// 原发票号码(开红票时传入)
	_normalInvoiceNo string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 消费者地址
	_payerAddress string
	// 付款方开票开户银行及账号
	_payerBankaccount string
	// 消费者电子邮箱
	_payerEmail string
	// 付款方名称, 对应发票台头
	_payerName string
	// 消费者联系电话
	_payerPhone string
	// 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
	_sumPrice string
	// 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
	_sumTax string
	// 复核人
	_payeeChecker string
	// 收款人
	_payeeReceiver string
	// 收款方电话
	_payeePhone string
	// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
	_applyId string
	// 发票种类，0=电子发票,1=纸质发票,2=专票
	_invoiceKind int64
	// 红字通知单号，冲红时需要，商家跟税局申请
	_redNoticeNo string
}

// New
