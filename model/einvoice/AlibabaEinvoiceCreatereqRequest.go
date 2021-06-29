package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP开票请求接口 API请求
alibaba.einvoice.createreq

ERP发起开票请求
*/
type AlibabaEinvoiceCreatereqRequest struct {
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
    // 付款方税务登记证号。对企业开具电子发票时必填。
    payerRegisterNo   string
    // 开票人
    payeeOperator   string
    // 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    invoiceAmount   string
    // 电子发票明细
    invoiceItems   []InvoiceItem
    // 发票备注，有些省市会把此信息打印到PDF中
    invoiceMemo   string
    // 开票日期, 格式"YYYY-MM-DD HH:SS:MM"
    invoiceTime   string
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
    // 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
    outShopName   string
    // 发票种类，0=电子发票,1=纸质发票,2=专票。注意：未订购纸票服务的税号无法开具纸票
    invoiceKind   int64
    // 红字通知单号，专票冲红时需要，商家跟税局申请
    redNoticeNo   string
    // 开票角色，supplier=供应商，只有platform_code=TB|TM时生效，供应商开票时允许使用A店铺开具B店铺的订单号
    requestRole   string
    // 分销订单号，request_role=supplier供应商开票时必填，分销订单号必须属于platform_tid，同时分销订单号的供应商必须和开票的授权账号一致
    distributionTid   string
}

// 初始化AlibabaEinvoiceCreatereqRequest对象
func NewAlibabaEinvoiceCreatereqRequest() *AlibabaEinvoiceCreatereqRequest{
    return &AlibabaEinvoiceCreatereqRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCreatereqRequest) GetApiMethodName() string {
    return "alibaba.einvoice.createreq"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCreatereqRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BusinessType Setter
// 默认：0。对于商家对个人开具，为0;对于商家对企业开具，为1;
func (r *AlibabaEinvoiceCreatereqRequest) SetBusinessType(businessType int64) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceCreatereqRequest) GetBusinessType() int64 {
    return r.businessType
}
// ErpTid Setter
// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
func (r *AlibabaEinvoiceCreatereqRequest) SetErpTid(erpTid string) error {
    r.erpTid = erpTid
    r.Set("erp_tid", erpTid)
    return nil
}

// ErpTid Getter
func (r AlibabaEinvoiceCreatereqRequest) GetErpTid() string {
    return r.erpTid
}
// PlatformCode Setter
// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
func (r *AlibabaEinvoiceCreatereqRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPlatformCode() string {
    return r.platformCode
}
// PlatformTid Setter
// 电商平台对应的主订单号
func (r *AlibabaEinvoiceCreatereqRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPlatformTid() string {
    return r.platformTid
}
// SerialNo Setter
// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
func (r *AlibabaEinvoiceCreatereqRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetSerialNo() string {
    return r.serialNo
}
// PayeeAddress Setter
// 开票方地址(新版中为必传)
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeAddress(payeeAddress string) error {
    r.payeeAddress = payeeAddress
    r.Set("payee_address", payeeAddress)
    return nil
}

// PayeeAddress Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeAddress() string {
    return r.payeeAddress
}
// PayeeBankaccount Setter
// 开票方银行及 帐号
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeBankaccount(payeeBankaccount string) error {
    r.payeeBankaccount = payeeBankaccount
    r.Set("payee_bankaccount", payeeBankaccount)
    return nil
}

// PayeeBankaccount Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeBankaccount() string {
    return r.payeeBankaccount
}
// PayeeName Setter
// 开票方名称，公司名(如:XX商城)
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeName() string {
    return r.payeeName
}
// PayerRegisterNo Setter
// 付款方税务登记证号。对企业开具电子发票时必填。
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerRegisterNo(payerRegisterNo string) error {
    r.payerRegisterNo = payerRegisterNo
    r.Set("payer_register_no", payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerRegisterNo() string {
    return r.payerRegisterNo
}
// PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeOperator(payeeOperator string) error {
    r.payeeOperator = payeeOperator
    r.Set("payee_operator", payeeOperator)
    return nil
}

// PayeeOperator Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeOperator() string {
    return r.payeeOperator
}
// InvoiceAmount Setter
// 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceAmount(invoiceAmount string) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceAmount() string {
    return r.invoiceAmount
}
// InvoiceItems Setter
// 电子发票明细
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceItems(invoiceItems []InvoiceItem) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceItems() []InvoiceItem {
    return r.invoiceItems
}
// InvoiceMemo Setter
// 发票备注，有些省市会把此信息打印到PDF中
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceMemo(invoiceMemo string) error {
    r.invoiceMemo = invoiceMemo
    r.Set("invoice_memo", invoiceMemo)
    return nil
}

// InvoiceMemo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceMemo() string {
    return r.invoiceMemo
}
// InvoiceTime Setter
// 开票日期, 格式"YYYY-MM-DD HH:SS:MM"
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceTime(invoiceTime string) error {
    r.invoiceTime = invoiceTime
    r.Set("invoice_time", invoiceTime)
    return nil
}

// InvoiceTime Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceTime() string {
    return r.invoiceTime
}
// InvoiceType Setter
// 发票(开票)类型，蓝票blue,红票red，默认blue
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceType(invoiceType string) error {
    r.invoiceType = invoiceType
    r.Set("invoice_type", invoiceType)
    return nil
}

// InvoiceType Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceType() string {
    return r.invoiceType
}
// NormalInvoiceCode Setter
// 原发票代码(开红票时传入)
func (r *AlibabaEinvoiceCreatereqRequest) SetNormalInvoiceCode(normalInvoiceCode string) error {
    r.normalInvoiceCode = normalInvoiceCode
    r.Set("normal_invoice_code", normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoiceCreatereqRequest) GetNormalInvoiceCode() string {
    return r.normalInvoiceCode
}
// NormalInvoiceNo Setter
// 原发票号码(开红票时传入)
func (r *AlibabaEinvoiceCreatereqRequest) SetNormalInvoiceNo(normalInvoiceNo string) error {
    r.normalInvoiceNo = normalInvoiceNo
    r.Set("normal_invoice_no", normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetNormalInvoiceNo() string {
    return r.normalInvoiceNo
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// PayerAddress Setter
// 消费者地址
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerAddress(payerAddress string) error {
    r.payerAddress = payerAddress
    r.Set("payer_address", payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerAddress() string {
    return r.payerAddress
}
// PayerBankaccount Setter
// 付款方开票开户银行及账号
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerBankaccount(payerBankaccount string) error {
    r.payerBankaccount = payerBankaccount
    r.Set("payer_bankaccount", payerBankaccount)
    return nil
}

// PayerBankaccount Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerBankaccount() string {
    return r.payerBankaccount
}
// PayerEmail Setter
// 消费者电子邮箱
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerEmail(payerEmail string) error {
    r.payerEmail = payerEmail
    r.Set("payer_email", payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerEmail() string {
    return r.payerEmail
}
// PayerName Setter
// 付款方名称, 对应发票台头
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerName(payerName string) error {
    r.payerName = payerName
    r.Set("payer_name", payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerName() string {
    return r.payerName
}
// PayerPhone Setter
// 消费者联系电话
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerPhone(payerPhone string) error {
    r.payerPhone = payerPhone
    r.Set("payer_phone", payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerPhone() string {
    return r.payerPhone
}
// SumPrice Setter
// 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceCreatereqRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceCreatereqRequest) GetSumPrice() string {
    return r.sumPrice
}
// SumTax Setter
// 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceCreatereqRequest) SetSumTax(sumTax string) error {
    r.sumTax = sumTax
    r.Set("sum_tax", sumTax)
    return nil
}

// SumTax Getter
func (r AlibabaEinvoiceCreatereqRequest) GetSumTax() string {
    return r.sumTax
}
// PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeChecker(payeeChecker string) error {
    r.payeeChecker = payeeChecker
    r.Set("payee_checker", payeeChecker)
    return nil
}

// PayeeChecker Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeChecker() string {
    return r.payeeChecker
}
// PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeReceiver(payeeReceiver string) error {
    r.payeeReceiver = payeeReceiver
    r.Set("payee_receiver", payeeReceiver)
    return nil
}

// PayeeReceiver Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeReceiver() string {
    return r.payeeReceiver
}
// PayeePhone Setter
// 收款方电话
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeePhone(payeePhone string) error {
    r.payeePhone = payeePhone
    r.Set("payee_phone", payeePhone)
    return nil
}

// PayeePhone Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeePhone() string {
    return r.payeePhone
}
// ApplyId Setter
// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
func (r *AlibabaEinvoiceCreatereqRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaEinvoiceCreatereqRequest) GetApplyId() string {
    return r.applyId
}
// OutShopName Setter
// 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
func (r *AlibabaEinvoiceCreatereqRequest) SetOutShopName(outShopName string) error {
    r.outShopName = outShopName
    r.Set("out_shop_name", outShopName)
    return nil
}

// OutShopName Getter
func (r AlibabaEinvoiceCreatereqRequest) GetOutShopName() string {
    return r.outShopName
}
// InvoiceKind Setter
// 发票种类，0=电子发票,1=纸质发票,2=专票。注意：未订购纸票服务的税号无法开具纸票
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}
// RedNoticeNo Setter
// 红字通知单号，专票冲红时需要，商家跟税局申请
func (r *AlibabaEinvoiceCreatereqRequest) SetRedNoticeNo(redNoticeNo string) error {
    r.redNoticeNo = redNoticeNo
    r.Set("red_notice_no", redNoticeNo)
    return nil
}

// RedNoticeNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetRedNoticeNo() string {
    return r.redNoticeNo
}
// RequestRole Setter
// 开票角色，supplier=供应商，只有platform_code=TB|TM时生效，供应商开票时允许使用A店铺开具B店铺的订单号
func (r *AlibabaEinvoiceCreatereqRequest) SetRequestRole(requestRole string) error {
    r.requestRole = requestRole
    r.Set("request_role", requestRole)
    return nil
}

// RequestRole Getter
func (r AlibabaEinvoiceCreatereqRequest) GetRequestRole() string {
    return r.requestRole
}
// DistributionTid Setter
// 分销订单号，request_role=supplier供应商开票时必填，分销订单号必须属于platform_tid，同时分销订单号的供应商必须和开票的授权账号一致
func (r *AlibabaEinvoiceCreatereqRequest) SetDistributionTid(distributionTid string) error {
    r.distributionTid = distributionTid
    r.Set("distribution_tid", distributionTid)
    return nil
}

// DistributionTid Getter
func (r AlibabaEinvoiceCreatereqRequest) GetDistributionTid() string {
    return r.distributionTid
}
