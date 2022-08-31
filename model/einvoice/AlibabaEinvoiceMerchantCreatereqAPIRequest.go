package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceMerchantCreatereqAPIRequest 商家自研ERP开票请求接口 API请求
// alibaba.einvoice.merchant.createreq
//
// 商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
type AlibabaEinvoiceMerchantCreatereqAPIRequest struct {
	model.Params
	// 电子发票明细
	_invoiceItems []InvoiceItem
	// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
	_platformCode string
	// 开票方银行及 帐号
	_payeeBankaccount string
	// 收款人
	_payeeReceiver string
	// 开票人
	_payeeOperator string
	// 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
	_invoiceAmount string
	// 复核人
	_payeeChecker string
	// 消费者联系电话
	_payerPhone string
	// 付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
	_payerRegisterNo string
	// 收款方电话
	_payeePhone string
	// 消费者电子邮箱
	_payerEmail string
	// 开票方名称，公司名(如:XX商城)
	_payeeName string
	// 消费者地址
	_payerAddress string
	// 发票备注，有些省市会把此信息打印到PDF中
	_invoiceMemo string
	// 付款方开票开户银行及账号
	_payerBankaccount string
	// 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
	_sumPrice string
	// 原发票号码(开红票时传入)
	_normalInvoiceNo string
	// 发票(开票)类型，蓝票blue,红票red，默认blue
	_invoiceType string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 原发票代码(开红票时传入)
	_normalInvoiceCode string
	// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
	_erpTid string
	// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
	_serialNo string
	// 电商平台对应的主订单号
	_platformTid string
	// 开票方地址(新版中为必传)
	_payeeAddress string
	// 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
	_sumTax string
	// 付款方名称, 对应发票台头
	_payerName string
	// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
	_applyId string
	// 红字通知单号，冲红时需要，商家跟税局申请
	_redNoticeNo string
	// 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
	_businessType int64
	// 发票种类，0=电子发票,1=纸质发票,2=专票
	_invoiceKind int64
}

// NewAlibabaEinvoiceMerchantCreatereqRequest 初始化AlibabaEinvoiceMerchantCreatereqAPIRequest对象
func NewAlibabaEinvoiceMerchantCreatereqRequest() *AlibabaEinvoiceMerchantCreatereqAPIRequest {
	return &AlibabaEinvoiceMerchantCreatereqAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.merchant.createreq"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInvoiceItems is InvoiceItems Setter
// 电子发票明细
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetInvoiceItems(_invoiceItems []InvoiceItem) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetInvoiceItems() []InvoiceItem {
	return r._invoiceItems
}

// SetPlatformCode is PlatformCode Setter
// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPayeeBankaccount is PayeeBankaccount Setter
// 开票方银行及 帐号
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeBankaccount(_payeeBankaccount string) error {
	r._payeeBankaccount = _payeeBankaccount
	r.Set("payee_bankaccount", _payeeBankaccount)
	return nil
}

// GetPayeeBankaccount PayeeBankaccount Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeBankaccount() string {
	return r._payeeBankaccount
}

// SetPayeeReceiver is PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeReceiver(_payeeReceiver string) error {
	r._payeeReceiver = _payeeReceiver
	r.Set("payee_receiver", _payeeReceiver)
	return nil
}

// GetPayeeReceiver PayeeReceiver Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeReceiver() string {
	return r._payeeReceiver
}

// SetPayeeOperator is PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeOperator(_payeeOperator string) error {
	r._payeeOperator = _payeeOperator
	r.Set("payee_operator", _payeeOperator)
	return nil
}

// GetPayeeOperator PayeeOperator Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeOperator() string {
	return r._payeeOperator
}

// SetInvoiceAmount is InvoiceAmount Setter
// 开票金额； &lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;当开红票时，该字段为负数&lt;/span&gt;
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetInvoiceAmount(_invoiceAmount string) error {
	r._invoiceAmount = _invoiceAmount
	r.Set("invoice_amount", _invoiceAmount)
	return nil
}

// GetInvoiceAmount InvoiceAmount Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetInvoiceAmount() string {
	return r._invoiceAmount
}

// SetPayeeChecker is PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeChecker(_payeeChecker string) error {
	r._payeeChecker = _payeeChecker
	r.Set("payee_checker", _payeeChecker)
	return nil
}

// GetPayeeChecker PayeeChecker Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeChecker() string {
	return r._payeeChecker
}

// SetPayerPhone is PayerPhone Setter
// 消费者联系电话
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayerPhone(_payerPhone string) error {
	r._payerPhone = _payerPhone
	r.Set("payer_phone", _payerPhone)
	return nil
}

// GetPayerPhone PayerPhone Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayerPhone() string {
	return r._payerPhone
}

// SetPayerRegisterNo is PayerRegisterNo Setter
// 付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
	r._payerRegisterNo = _payerRegisterNo
	r.Set("payer_register_no", _payerRegisterNo)
	return nil
}

// GetPayerRegisterNo PayerRegisterNo Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayerRegisterNo() string {
	return r._payerRegisterNo
}

// SetPayeePhone is PayeePhone Setter
// 收款方电话
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeePhone(_payeePhone string) error {
	r._payeePhone = _payeePhone
	r.Set("payee_phone", _payeePhone)
	return nil
}

// GetPayeePhone PayeePhone Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeePhone() string {
	return r._payeePhone
}

// SetPayerEmail is PayerEmail Setter
// 消费者电子邮箱
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayerEmail(_payerEmail string) error {
	r._payerEmail = _payerEmail
	r.Set("payer_email", _payerEmail)
	return nil
}

// GetPayerEmail PayerEmail Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayerEmail() string {
	return r._payerEmail
}

// SetPayeeName is PayeeName Setter
// 开票方名称，公司名(如:XX商城)
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// GetPayeeName PayeeName Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// SetPayerAddress is PayerAddress Setter
// 消费者地址
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayerAddress(_payerAddress string) error {
	r._payerAddress = _payerAddress
	r.Set("payer_address", _payerAddress)
	return nil
}

// GetPayerAddress PayerAddress Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayerAddress() string {
	return r._payerAddress
}

// SetInvoiceMemo is InvoiceMemo Setter
// 发票备注，有些省市会把此信息打印到PDF中
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetInvoiceMemo(_invoiceMemo string) error {
	r._invoiceMemo = _invoiceMemo
	r.Set("invoice_memo", _invoiceMemo)
	return nil
}

// GetInvoiceMemo InvoiceMemo Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetInvoiceMemo() string {
	return r._invoiceMemo
}

// SetPayerBankaccount is PayerBankaccount Setter
// 付款方开票开户银行及账号
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayerBankaccount(_payerBankaccount string) error {
	r._payerBankaccount = _payerBankaccount
	r.Set("payer_bankaccount", _payerBankaccount)
	return nil
}

// GetPayerBankaccount PayerBankaccount Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayerBankaccount() string {
	return r._payerBankaccount
}

// SetSumPrice is SumPrice Setter
// 合计金额(新版中为必传) &lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;当开红票时，该字段为负数&lt;/span&gt;
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetNormalInvoiceNo is NormalInvoiceNo Setter
// 原发票号码(开红票时传入)
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
	r._normalInvoiceNo = _normalInvoiceNo
	r.Set("normal_invoice_no", _normalInvoiceNo)
	return nil
}

// GetNormalInvoiceNo NormalInvoiceNo Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetNormalInvoiceNo() string {
	return r._normalInvoiceNo
}

// SetInvoiceType is InvoiceType Setter
// 发票(开票)类型，蓝票blue,红票red，默认blue
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetInvoiceType(_invoiceType string) error {
	r._invoiceType = _invoiceType
	r.Set("invoice_type", _invoiceType)
	return nil
}

// GetInvoiceType InvoiceType Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetInvoiceType() string {
	return r._invoiceType
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetNormalInvoiceCode is NormalInvoiceCode Setter
// 原发票代码(开红票时传入)
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
	r._normalInvoiceCode = _normalInvoiceCode
	r.Set("normal_invoice_code", _normalInvoiceCode)
	return nil
}

// GetNormalInvoiceCode NormalInvoiceCode Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetNormalInvoiceCode() string {
	return r._normalInvoiceCode
}

// SetErpTid is ErpTid Setter
// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetErpTid(_erpTid string) error {
	r._erpTid = _erpTid
	r.Set("erp_tid", _erpTid)
	return nil
}

// GetErpTid ErpTid Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetErpTid() string {
	return r._erpTid
}

// SetSerialNo is SerialNo Setter
// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetPlatformTid is PlatformTid Setter
// 电商平台对应的主订单号
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetPayeeAddress is PayeeAddress Setter
// 开票方地址(新版中为必传)
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayeeAddress(_payeeAddress string) error {
	r._payeeAddress = _payeeAddress
	r.Set("payee_address", _payeeAddress)
	return nil
}

// GetPayeeAddress PayeeAddress Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayeeAddress() string {
	return r._payeeAddress
}

// SetSumTax is SumTax Setter
// 合计税额 &lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;当开红票时，该字段为负数&lt;/span&gt;
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetSumTax(_sumTax string) error {
	r._sumTax = _sumTax
	r.Set("sum_tax", _sumTax)
	return nil
}

// GetSumTax SumTax Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetSumTax() string {
	return r._sumTax
}

// SetPayerName is PayerName Setter
// 付款方名称, 对应发票台头
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetPayerName(_payerName string) error {
	r._payerName = _payerName
	r.Set("payer_name", _payerName)
	return nil
}

// GetPayerName PayerName Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetPayerName() string {
	return r._payerName
}

// SetApplyId is ApplyId Setter
// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetRedNoticeNo is RedNoticeNo Setter
// 红字通知单号，冲红时需要，商家跟税局申请
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetRedNoticeNo(_redNoticeNo string) error {
	r._redNoticeNo = _redNoticeNo
	r.Set("red_notice_no", _redNoticeNo)
	return nil
}

// GetRedNoticeNo RedNoticeNo Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetRedNoticeNo() string {
	return r._redNoticeNo
}

// SetBusinessType is BusinessType Setter
// 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetBusinessType() int64 {
	return r._businessType
}

// SetInvoiceKind is InvoiceKind Setter
// 发票种类，0=电子发票,1=纸质发票,2=专票
func (r *AlibabaEinvoiceMerchantCreatereqAPIRequest) SetInvoiceKind(_invoiceKind int64) error {
	r._invoiceKind = _invoiceKind
	r.Set("invoice_kind", _invoiceKind)
	return nil
}

// GetInvoiceKind InvoiceKind Getter
func (r AlibabaEinvoiceMerchantCreatereqAPIRequest) GetInvoiceKind() int64 {
	return r._invoiceKind
}
