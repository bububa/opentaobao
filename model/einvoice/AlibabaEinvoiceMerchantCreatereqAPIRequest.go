package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicemerchantcreatereqAPIRequest 商家自研ERP开票请求接口 API请求
// alibaba.einvoice.merchant.createreq
//
// 商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
type AlibabaeinvoicemerchantcreatereqAPIRequest struct {
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

// NewAlibabaeinvoicemerchantcreatereqRequest 初始化AlibabaeinvoicemerchantcreatereqAPIRequest对象
func NewAlibabaeinvoicemerchantcreatereqRequest() *AlibabaeinvoicemerchantcreatereqAPIRequest {
	return &AlibabaeinvoicemerchantcreatereqAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.merchant.createreq"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceItems is InvoiceItems Setter
// 电子发票明细
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetInvoiceItems(_invoiceItems []InvoiceItem) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetInvoiceItems() []InvoiceItem {
	return r._invoiceItems
}

// SetPlatformCode is PlatformCode Setter
// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPayeeBankaccount is PayeeBankaccount Setter
// 开票方银行及 帐号
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeBankaccount(_payeeBankaccount string) error {
	r._payeeBankaccount = _payeeBankaccount
	r.Set("payee_bankaccount", _payeeBankaccount)
	return nil
}

// GetPayeeBankaccount PayeeBankaccount Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeBankaccount() string {
	return r._payeeBankaccount
}

// SetPayeeReceiver is PayeeReceiver Setter
// 收款人
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeReceiver(_payeeReceiver string) error {
	r._payeeReceiver = _payeeReceiver
	r.Set("payee_receiver", _payeeReceiver)
	return nil
}

// GetPayeeReceiver PayeeReceiver Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeReceiver() string {
	return r._payeeReceiver
}

// SetPayeeOperator is PayeeOperator Setter
// 开票人
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeOperator(_payeeOperator string) error {
	r._payeeOperator = _payeeOperator
	r.Set("payee_operator", _payeeOperator)
	return nil
}

// GetPayeeOperator PayeeOperator Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeOperator() string {
	return r._payeeOperator
}

// SetInvoiceAmount is InvoiceAmount Setter
// 开票金额； &lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;当开红票时，该字段为负数&lt;/span&gt;
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetInvoiceAmount(_invoiceAmount string) error {
	r._invoiceAmount = _invoiceAmount
	r.Set("invoice_amount", _invoiceAmount)
	return nil
}

// GetInvoiceAmount InvoiceAmount Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetInvoiceAmount() string {
	return r._invoiceAmount
}

// SetPayeeChecker is PayeeChecker Setter
// 复核人
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeChecker(_payeeChecker string) error {
	r._payeeChecker = _payeeChecker
	r.Set("payee_checker", _payeeChecker)
	return nil
}

// GetPayeeChecker PayeeChecker Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeChecker() string {
	return r._payeeChecker
}

// SetPayerPhone is PayerPhone Setter
// 消费者联系电话
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayerPhone(_payerPhone string) error {
	r._payerPhone = _payerPhone
	r.Set("payer_phone", _payerPhone)
	return nil
}

// GetPayerPhone PayerPhone Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayerPhone() string {
	return r._payerPhone
}

// SetPayerRegisterNo is PayerRegisterNo Setter
// 付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
	r._payerRegisterNo = _payerRegisterNo
	r.Set("payer_register_no", _payerRegisterNo)
	return nil
}

// GetPayerRegisterNo PayerRegisterNo Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayerRegisterNo() string {
	return r._payerRegisterNo
}

// SetPayeePhone is PayeePhone Setter
// 收款方电话
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeePhone(_payeePhone string) error {
	r._payeePhone = _payeePhone
	r.Set("payee_phone", _payeePhone)
	return nil
}

// GetPayeePhone PayeePhone Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeePhone() string {
	return r._payeePhone
}

// SetPayerEmail is PayerEmail Setter
// 消费者电子邮箱
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayerEmail(_payerEmail string) error {
	r._payerEmail = _payerEmail
	r.Set("payer_email", _payerEmail)
	return nil
}

// GetPayerEmail PayerEmail Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayerEmail() string {
	return r._payerEmail
}

// SetPayeeName is PayeeName Setter
// 开票方名称，公司名(如:XX商城)
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// GetPayeeName PayeeName Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// SetPayerAddress is PayerAddress Setter
// 消费者地址
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayerAddress(_payerAddress string) error {
	r._payerAddress = _payerAddress
	r.Set("payer_address", _payerAddress)
	return nil
}

// GetPayerAddress PayerAddress Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayerAddress() string {
	return r._payerAddress
}

// SetInvoiceMemo is InvoiceMemo Setter
// 发票备注，有些省市会把此信息打印到PDF中
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetInvoiceMemo(_invoiceMemo string) error {
	r._invoiceMemo = _invoiceMemo
	r.Set("invoice_memo", _invoiceMemo)
	return nil
}

// GetInvoiceMemo InvoiceMemo Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetInvoiceMemo() string {
	return r._invoiceMemo
}

// SetPayerBankaccount is PayerBankaccount Setter
// 付款方开票开户银行及账号
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayerBankaccount(_payerBankaccount string) error {
	r._payerBankaccount = _payerBankaccount
	r.Set("payer_bankaccount", _payerBankaccount)
	return nil
}

// GetPayerBankaccount PayerBankaccount Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayerBankaccount() string {
	return r._payerBankaccount
}

// SetSumPrice is SumPrice Setter
// 合计金额(新版中为必传) &lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;当开红票时，该字段为负数&lt;/span&gt;
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetNormalInvoiceNo is NormalInvoiceNo Setter
// 原发票号码(开红票时传入)
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
	r._normalInvoiceNo = _normalInvoiceNo
	r.Set("normal_invoice_no", _normalInvoiceNo)
	return nil
}

// GetNormalInvoiceNo NormalInvoiceNo Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetNormalInvoiceNo() string {
	return r._normalInvoiceNo
}

// SetInvoiceType is InvoiceType Setter
// 发票(开票)类型，蓝票blue,红票red，默认blue
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetInvoiceType(_invoiceType string) error {
	r._invoiceType = _invoiceType
	r.Set("invoice_type", _invoiceType)
	return nil
}

// GetInvoiceType InvoiceType Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetInvoiceType() string {
	return r._invoiceType
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetNormalInvoiceCode is NormalInvoiceCode Setter
// 原发票代码(开红票时传入)
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
	r._normalInvoiceCode = _normalInvoiceCode
	r.Set("normal_invoice_code", _normalInvoiceCode)
	return nil
}

// GetNormalInvoiceCode NormalInvoiceCode Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetNormalInvoiceCode() string {
	return r._normalInvoiceCode
}

// SetErpTid is ErpTid Setter
// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetErpTid(_erpTid string) error {
	r._erpTid = _erpTid
	r.Set("erp_tid", _erpTid)
	return nil
}

// GetErpTid ErpTid Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetErpTid() string {
	return r._erpTid
}

// SetSerialNo is SerialNo Setter
// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetPlatformTid is PlatformTid Setter
// 电商平台对应的主订单号
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetPayeeAddress is PayeeAddress Setter
// 开票方地址(新版中为必传)
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayeeAddress(_payeeAddress string) error {
	r._payeeAddress = _payeeAddress
	r.Set("payee_address", _payeeAddress)
	return nil
}

// GetPayeeAddress PayeeAddress Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayeeAddress() string {
	return r._payeeAddress
}

// SetSumTax is SumTax Setter
// 合计税额 &lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;当开红票时，该字段为负数&lt;/span&gt;
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetSumTax(_sumTax string) error {
	r._sumTax = _sumTax
	r.Set("sum_tax", _sumTax)
	return nil
}

// GetSumTax SumTax Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetSumTax() string {
	return r._sumTax
}

// SetPayerName is PayerName Setter
// 付款方名称, 对应发票台头
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetPayerName(_payerName string) error {
	r._payerName = _payerName
	r.Set("payer_name", _payerName)
	return nil
}

// GetPayerName PayerName Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetPayerName() string {
	return r._payerName
}

// SetApplyId is ApplyId Setter
// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetRedNoticeNo is RedNoticeNo Setter
// 红字通知单号，冲红时需要，商家跟税局申请
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetRedNoticeNo(_redNoticeNo string) error {
	r._redNoticeNo = _redNoticeNo
	r.Set("red_notice_no", _redNoticeNo)
	return nil
}

// GetRedNoticeNo RedNoticeNo Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetRedNoticeNo() string {
	return r._redNoticeNo
}

// SetBusinessType is BusinessType Setter
// 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetBusinessType() int64 {
	return r._businessType
}

// SetInvoiceKind is InvoiceKind Setter
// 发票种类，0=电子发票,1=纸质发票,2=专票
func (r *AlibabaeinvoicemerchantcreatereqAPIRequest) SetInvoiceKind(_invoiceKind int64) error {
	r._invoiceKind = _invoiceKind
	r.Set("invoice_kind", _invoiceKind)
	return nil
}

// GetInvoiceKind InvoiceKind Getter
func (r AlibabaeinvoicemerchantcreatereqAPIRequest) GetInvoiceKind() int64 {
	return r._invoiceKind
}
