package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自研ERP开票请求接口 API请求
alibaba.einvoice.merchant.createreq

商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
*/
type AlibabaEinvoiceMerchantCreatereqRequest struct {
    model.Params
    // 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
    businessType   int64
    // ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
    erpTid   string
    // 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    platformCode   string
    // 电商平台对应的主订单号
    platformTid   string
    // 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
    serialNo   string
    // 开票方地址(新版中为必传)
    payeeAddress   string
    // 开票方银行及 帐号
    payeeBankaccount   string
    // 开票方名称，公司名(如:XX商城)
    payeeName   string
    // 付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
    payerRegisterNo   string
    // 开票人
    payeeOperator   string
    // 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    invoiceAmount   string
    // 电子发票明细
    invoiceItems   []InvoiceItem
    // 发票备注，有些省市会把此信息打印到PDF中
    invoiceMemo   string
    // 发票(开票)类型，蓝票blue,红票red，默认blue
    invoiceType   string
    // 原发票代码(开红票时传入)
    normalInvoiceCode   string
    // 原发票号码(开红票时传入)
    normalInvoiceNo   string
    // 收款方税务登记证号
    payeeRegisterNo   string
    // 消费者地址
    payerAddress   string
    // 付款方开票开户银行及账号
    payerBankaccount   string
    // 消费者电子邮箱
    payerEmail   string
    // 付款方名称, 对应发票台头
    payerName   string
    // 消费者联系电话
    payerPhone   string
    // 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    sumPrice   string
    // 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    sumTax   string
    // 复核人
    payeeChecker   string
    // 收款人
    payeeReceiver   string
    // 收款方电话
    payeePhone   string
    // 开票申请ID，接收了开票申请消息后，需要把apply_id带上
    applyId   string
    // 发票种类，0=电子发票,1=纸质发票,2=专票
    invoiceKind   int64
    // 红字通知单号，冲红时需要，商家跟税局申请
    redNoticeNo   string
}

// 初始化AlibabaEinvoiceMerchantCreatereqRequest对象
func NewAlibabaEinvoiceMerchantCreatereqRequest() *AlibabaEinvoiceMerchantCreatereqRequest{
    return &AlibabaEinvoiceMerchantCreatereqRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetApiMethodName() string {
    return "alibaba.einvoice.merchant.createreq"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BusinessType Setter
// 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetBusinessType(businessType int64) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetBusinessType() int64 {
    return r.businessType
}
// ErpTid Setter
// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetErpTid(erpTid string) error {
    r.erpTid = erpTid
    r.Set("erp_tid", erpTid)
    return nil
}

// ErpTid Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetErpTid() string {
    return r.erpTid
}
// PlatformCode Setter
// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPlatformCode() string {
    return r.platformCode
}
// PlatformTid Setter
// 电商平台对应的主订单号
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPlatformTid() string {
    return r.platformTid
}
// SerialNo Setter
// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetSerialNo() string {
    return r.serialNo
}
// PayeeAddress Setter
// 开票方地址(新版中为必传)
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeAddress(payeeAddress string) error {
    r.payeeAddress = payeeAddress
    r.Set("payee_address", payeeAddress)
    return nil
}

// PayeeAddress Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeAddress() string {
    return r.payeeAddress
}
// PayeeBankaccount Setter
// 开票方银行及 帐号
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeBankaccount(payeeBankaccount string) error {
    r.payeeBankaccount = payeeBankaccount
    r.Set("payee_bankaccount", payeeBankaccount)
    return nil
}

// PayeeBankaccount Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeBankaccount() string {
    return r.payeeBankaccount
}
// PayeeName Setter
// 开票方名称，公司名(如:XX商城)
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeName() string {
    return r.payeeName
}
// PayerRegisterNo Setter
// 付款方税务登记证号。对企业开具电子发票时必填。目前北京地区暂未开放对企业开具电子发票，若北京地区放开后，对于向企业开具的情况，付款方税务登记证号和名称也不能为空
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayerRegisterNo(payerRegisterNo string) error {
    r.payerRegisterNo = payerRegisterNo
    r.Set("payer_register_no", payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayerRegisterNo() string {
    return r.payerRegisterNo
}
// PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeOperator(payeeOperator string) error {
    r.payeeOperator = payeeOperator
    r.Set("payee_operator", payeeOperator)
    return nil
}

// PayeeOperator Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeOperator() string {
    return r.payeeOperator
}
// InvoiceAmount Setter
// 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetInvoiceAmount(invoiceAmount string) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetInvoiceAmount() string {
    return r.invoiceAmount
}
// InvoiceItems Setter
// 电子发票明细
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetInvoiceItems(invoiceItems []InvoiceItem) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetInvoiceItems() []InvoiceItem {
    return r.invoiceItems
}
// InvoiceMemo Setter
// 发票备注，有些省市会把此信息打印到PDF中
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetInvoiceMemo(invoiceMemo string) error {
    r.invoiceMemo = invoiceMemo
    r.Set("invoice_memo", invoiceMemo)
    return nil
}

// InvoiceMemo Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetInvoiceMemo() string {
    return r.invoiceMemo
}
// InvoiceType Setter
// 发票(开票)类型，蓝票blue,红票red，默认blue
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetInvoiceType(invoiceType string) error {
    r.invoiceType = invoiceType
    r.Set("invoice_type", invoiceType)
    return nil
}

// InvoiceType Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetInvoiceType() string {
    return r.invoiceType
}
// NormalInvoiceCode Setter
// 原发票代码(开红票时传入)
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetNormalInvoiceCode(normalInvoiceCode string) error {
    r.normalInvoiceCode = normalInvoiceCode
    r.Set("normal_invoice_code", normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetNormalInvoiceCode() string {
    return r.normalInvoiceCode
}
// NormalInvoiceNo Setter
// 原发票号码(开红票时传入)
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetNormalInvoiceNo(normalInvoiceNo string) error {
    r.normalInvoiceNo = normalInvoiceNo
    r.Set("normal_invoice_no", normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetNormalInvoiceNo() string {
    return r.normalInvoiceNo
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// PayerAddress Setter
// 消费者地址
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayerAddress(payerAddress string) error {
    r.payerAddress = payerAddress
    r.Set("payer_address", payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayerAddress() string {
    return r.payerAddress
}
// PayerBankaccount Setter
// 付款方开票开户银行及账号
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayerBankaccount(payerBankaccount string) error {
    r.payerBankaccount = payerBankaccount
    r.Set("payer_bankaccount", payerBankaccount)
    return nil
}

// PayerBankaccount Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayerBankaccount() string {
    return r.payerBankaccount
}
// PayerEmail Setter
// 消费者电子邮箱
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayerEmail(payerEmail string) error {
    r.payerEmail = payerEmail
    r.Set("payer_email", payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayerEmail() string {
    return r.payerEmail
}
// PayerName Setter
// 付款方名称, 对应发票台头
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayerName(payerName string) error {
    r.payerName = payerName
    r.Set("payer_name", payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayerName() string {
    return r.payerName
}
// PayerPhone Setter
// 消费者联系电话
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayerPhone(payerPhone string) error {
    r.payerPhone = payerPhone
    r.Set("payer_phone", payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayerPhone() string {
    return r.payerPhone
}
// SumPrice Setter
// 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetSumPrice() string {
    return r.sumPrice
}
// SumTax Setter
// 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetSumTax(sumTax string) error {
    r.sumTax = sumTax
    r.Set("sum_tax", sumTax)
    return nil
}

// SumTax Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetSumTax() string {
    return r.sumTax
}
// PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeChecker(payeeChecker string) error {
    r.payeeChecker = payeeChecker
    r.Set("payee_checker", payeeChecker)
    return nil
}

// PayeeChecker Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeChecker() string {
    return r.payeeChecker
}
// PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeeReceiver(payeeReceiver string) error {
    r.payeeReceiver = payeeReceiver
    r.Set("payee_receiver", payeeReceiver)
    return nil
}

// PayeeReceiver Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeeReceiver() string {
    return r.payeeReceiver
}
// PayeePhone Setter
// 收款方电话
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetPayeePhone(payeePhone string) error {
    r.payeePhone = payeePhone
    r.Set("payee_phone", payeePhone)
    return nil
}

// PayeePhone Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetPayeePhone() string {
    return r.payeePhone
}
// ApplyId Setter
// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetApplyId() string {
    return r.applyId
}
// InvoiceKind Setter
// 发票种类，0=电子发票,1=纸质发票,2=专票
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}
// RedNoticeNo Setter
// 红字通知单号，冲红时需要，商家跟税局申请
func (r *AlibabaEinvoiceMerchantCreatereqRequest) SetRedNoticeNo(redNoticeNo string) error {
    r.redNoticeNo = redNoticeNo
    r.Set("red_notice_no", redNoticeNo)
    return nil
}

// RedNoticeNo Getter
func (r AlibabaEinvoiceMerchantCreatereqRequest) GetRedNoticeNo() string {
    return r.redNoticeNo
}
