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
    _businessType   int64
    // ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
    _erpTid   string
    // 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    _platformCode   string
    // 电商平台对应的主订单号
    _platformTid   string
    // 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
    _serialNo   string
    // 开票方地址(新版中为必传)
    _payeeAddress   string
    // 开票方银行及 帐号
    _payeeBankaccount   string
    // 开票方名称，公司名(如:XX商城)
    _payeeName   string
    // 付款方税务登记证号。对企业开具电子发票时必填。
    _payerRegisterNo   string
    // 开票人
    _payeeOperator   string
    // 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    _invoiceAmount   string
    // 电子发票明细
    _invoiceItems   []InvoiceItem
    // 发票备注，有些省市会把此信息打印到PDF中
    _invoiceMemo   string
    // 开票日期, 格式"YYYY-MM-DD HH:SS:MM"
    _invoiceTime   string
    // 发票(开票)类型，蓝票blue,红票red，默认blue
    _invoiceType   string
    // 原发票代码(开红票时传入)
    _normalInvoiceCode   string
    // 原发票号码(开红票时传入)
    _normalInvoiceNo   string
    // 收款方税务登记证号
    _payeeRegisterNo   string
    // 消费者地址
    _payerAddress   string
    // 付款方开票开户银行及账号
    _payerBankaccount   string
    // 消费者电子邮箱
    _payerEmail   string
    // 付款方名称, 对应发票台头
    _payerName   string
    // 消费者联系电话
    _payerPhone   string
    // 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    _sumPrice   string
    // 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
    _sumTax   string
    // 复核人
    _payeeChecker   string
    // 收款人
    _payeeReceiver   string
    // 收款方电话
    _payeePhone   string
    // 开票申请ID，接收了开票申请消息后，需要把apply_id带上
    _applyId   string
    // 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
    _outShopName   string
    // 发票种类，0=电子发票,1=纸质发票,2=专票。注意：未订购纸票服务的税号无法开具纸票
    _invoiceKind   int64
    // 红字通知单号，专票冲红时需要，商家跟税局申请
    _redNoticeNo   string
    // 开票角色，supplier=供应商，只有platform_code=TB|TM时生效，供应商开票时允许使用A店铺开具B店铺的订单号
    _requestRole   string
    // 分销订单号，request_role=supplier供应商开票时必填，分销订单号必须属于platform_tid，同时分销订单号的供应商必须和开票的授权账号一致
    _distributionTid   string
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
func (r *AlibabaEinvoiceCreatereqRequest) SetBusinessType(_businessType int64) error {
    r._businessType = _businessType
    r.Set("business_type", _businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceCreatereqRequest) GetBusinessType() int64 {
    return r._businessType
}
// ErpTid Setter
// ERP系统中的单据号。如果没有erp的唯一单据号。建议使用platform_code+”_”+ platform_tid的组合方式
func (r *AlibabaEinvoiceCreatereqRequest) SetErpTid(_erpTid string) error {
    r._erpTid = _erpTid
    r.Set("erp_tid", _erpTid)
    return nil
}

// ErpTid Getter
func (r AlibabaEinvoiceCreatereqRequest) GetErpTid() string {
    return r._erpTid
}
// PlatformCode Setter
// 电商平台代码。TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
func (r *AlibabaEinvoiceCreatereqRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPlatformCode() string {
    return r._platformCode
}
// PlatformTid Setter
// 电商平台对应的主订单号
func (r *AlibabaEinvoiceCreatereqRequest) SetPlatformTid(_platformTid string) error {
    r._platformTid = _platformTid
    r.Set("platform_tid", _platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPlatformTid() string {
    return r._platformTid
}
// SerialNo Setter
// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。请调用平台统一流水号获取接口，alibaba.einvoice.serialno.generate。
func (r *AlibabaEinvoiceCreatereqRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetSerialNo() string {
    return r._serialNo
}
// PayeeAddress Setter
// 开票方地址(新版中为必传)
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeAddress(_payeeAddress string) error {
    r._payeeAddress = _payeeAddress
    r.Set("payee_address", _payeeAddress)
    return nil
}

// PayeeAddress Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeAddress() string {
    return r._payeeAddress
}
// PayeeBankaccount Setter
// 开票方银行及 帐号
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeBankaccount(_payeeBankaccount string) error {
    r._payeeBankaccount = _payeeBankaccount
    r.Set("payee_bankaccount", _payeeBankaccount)
    return nil
}

// PayeeBankaccount Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeBankaccount() string {
    return r._payeeBankaccount
}
// PayeeName Setter
// 开票方名称，公司名(如:XX商城)
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeName(_payeeName string) error {
    r._payeeName = _payeeName
    r.Set("payee_name", _payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeName() string {
    return r._payeeName
}
// PayerRegisterNo Setter
// 付款方税务登记证号。对企业开具电子发票时必填。
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
    r._payerRegisterNo = _payerRegisterNo
    r.Set("payer_register_no", _payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerRegisterNo() string {
    return r._payerRegisterNo
}
// PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeOperator(_payeeOperator string) error {
    r._payeeOperator = _payeeOperator
    r.Set("payee_operator", _payeeOperator)
    return nil
}

// PayeeOperator Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeOperator() string {
    return r._payeeOperator
}
// InvoiceAmount Setter
// 开票金额； <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceAmount(_invoiceAmount string) error {
    r._invoiceAmount = _invoiceAmount
    r.Set("invoice_amount", _invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceAmount() string {
    return r._invoiceAmount
}
// InvoiceItems Setter
// 电子发票明细
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceItems(_invoiceItems []InvoiceItem) error {
    r._invoiceItems = _invoiceItems
    r.Set("invoice_items", _invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceItems() []InvoiceItem {
    return r._invoiceItems
}
// InvoiceMemo Setter
// 发票备注，有些省市会把此信息打印到PDF中
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceMemo(_invoiceMemo string) error {
    r._invoiceMemo = _invoiceMemo
    r.Set("invoice_memo", _invoiceMemo)
    return nil
}

// InvoiceMemo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceMemo() string {
    return r._invoiceMemo
}
// InvoiceTime Setter
// 开票日期, 格式"YYYY-MM-DD HH:SS:MM"
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceTime(_invoiceTime string) error {
    r._invoiceTime = _invoiceTime
    r.Set("invoice_time", _invoiceTime)
    return nil
}

// InvoiceTime Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceTime() string {
    return r._invoiceTime
}
// InvoiceType Setter
// 发票(开票)类型，蓝票blue,红票red，默认blue
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceType(_invoiceType string) error {
    r._invoiceType = _invoiceType
    r.Set("invoice_type", _invoiceType)
    return nil
}

// InvoiceType Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceType() string {
    return r._invoiceType
}
// NormalInvoiceCode Setter
// 原发票代码(开红票时传入)
func (r *AlibabaEinvoiceCreatereqRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
    r._normalInvoiceCode = _normalInvoiceCode
    r.Set("normal_invoice_code", _normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoiceCreatereqRequest) GetNormalInvoiceCode() string {
    return r._normalInvoiceCode
}
// NormalInvoiceNo Setter
// 原发票号码(开红票时传入)
func (r *AlibabaEinvoiceCreatereqRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
    r._normalInvoiceNo = _normalInvoiceNo
    r.Set("normal_invoice_no", _normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetNormalInvoiceNo() string {
    return r._normalInvoiceNo
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// PayerAddress Setter
// 消费者地址
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerAddress(_payerAddress string) error {
    r._payerAddress = _payerAddress
    r.Set("payer_address", _payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerAddress() string {
    return r._payerAddress
}
// PayerBankaccount Setter
// 付款方开票开户银行及账号
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerBankaccount(_payerBankaccount string) error {
    r._payerBankaccount = _payerBankaccount
    r.Set("payer_bankaccount", _payerBankaccount)
    return nil
}

// PayerBankaccount Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerBankaccount() string {
    return r._payerBankaccount
}
// PayerEmail Setter
// 消费者电子邮箱
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerEmail(_payerEmail string) error {
    r._payerEmail = _payerEmail
    r.Set("payer_email", _payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerEmail() string {
    return r._payerEmail
}
// PayerName Setter
// 付款方名称, 对应发票台头
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerName(_payerName string) error {
    r._payerName = _payerName
    r.Set("payer_name", _payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerName() string {
    return r._payerName
}
// PayerPhone Setter
// 消费者联系电话
func (r *AlibabaEinvoiceCreatereqRequest) SetPayerPhone(_payerPhone string) error {
    r._payerPhone = _payerPhone
    r.Set("payer_phone", _payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayerPhone() string {
    return r._payerPhone
}
// SumPrice Setter
// 合计金额(新版中为必传) <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceCreatereqRequest) SetSumPrice(_sumPrice string) error {
    r._sumPrice = _sumPrice
    r.Set("sum_price", _sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceCreatereqRequest) GetSumPrice() string {
    return r._sumPrice
}
// SumTax Setter
// 合计税额 <span style="color:red;font-weight: bold;">当开红票时，该字段为负数</span>
func (r *AlibabaEinvoiceCreatereqRequest) SetSumTax(_sumTax string) error {
    r._sumTax = _sumTax
    r.Set("sum_tax", _sumTax)
    return nil
}

// SumTax Getter
func (r AlibabaEinvoiceCreatereqRequest) GetSumTax() string {
    return r._sumTax
}
// PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeChecker(_payeeChecker string) error {
    r._payeeChecker = _payeeChecker
    r.Set("payee_checker", _payeeChecker)
    return nil
}

// PayeeChecker Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeChecker() string {
    return r._payeeChecker
}
// PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeeReceiver(_payeeReceiver string) error {
    r._payeeReceiver = _payeeReceiver
    r.Set("payee_receiver", _payeeReceiver)
    return nil
}

// PayeeReceiver Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeeReceiver() string {
    return r._payeeReceiver
}
// PayeePhone Setter
// 收款方电话
func (r *AlibabaEinvoiceCreatereqRequest) SetPayeePhone(_payeePhone string) error {
    r._payeePhone = _payeePhone
    r.Set("payee_phone", _payeePhone)
    return nil
}

// PayeePhone Getter
func (r AlibabaEinvoiceCreatereqRequest) GetPayeePhone() string {
    return r._payeePhone
}
// ApplyId Setter
// 开票申请ID，接收了开票申请消息后，需要把apply_id带上
func (r *AlibabaEinvoiceCreatereqRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaEinvoiceCreatereqRequest) GetApplyId() string {
    return r._applyId
}
// OutShopName Setter
// 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
func (r *AlibabaEinvoiceCreatereqRequest) SetOutShopName(_outShopName string) error {
    r._outShopName = _outShopName
    r.Set("out_shop_name", _outShopName)
    return nil
}

// OutShopName Getter
func (r AlibabaEinvoiceCreatereqRequest) GetOutShopName() string {
    return r._outShopName
}
// InvoiceKind Setter
// 发票种类，0=电子发票,1=纸质发票,2=专票。注意：未订购纸票服务的税号无法开具纸票
func (r *AlibabaEinvoiceCreatereqRequest) SetInvoiceKind(_invoiceKind int64) error {
    r._invoiceKind = _invoiceKind
    r.Set("invoice_kind", _invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceCreatereqRequest) GetInvoiceKind() int64 {
    return r._invoiceKind
}
// RedNoticeNo Setter
// 红字通知单号，专票冲红时需要，商家跟税局申请
func (r *AlibabaEinvoiceCreatereqRequest) SetRedNoticeNo(_redNoticeNo string) error {
    r._redNoticeNo = _redNoticeNo
    r.Set("red_notice_no", _redNoticeNo)
    return nil
}

// RedNoticeNo Getter
func (r AlibabaEinvoiceCreatereqRequest) GetRedNoticeNo() string {
    return r._redNoticeNo
}
// RequestRole Setter
// 开票角色，supplier=供应商，只有platform_code=TB|TM时生效，供应商开票时允许使用A店铺开具B店铺的订单号
func (r *AlibabaEinvoiceCreatereqRequest) SetRequestRole(_requestRole string) error {
    r._requestRole = _requestRole
    r.Set("request_role", _requestRole)
    return nil
}

// RequestRole Getter
func (r AlibabaEinvoiceCreatereqRequest) GetRequestRole() string {
    return r._requestRole
}
// DistributionTid Setter
// 分销订单号，request_role=supplier供应商开票时必填，分销订单号必须属于platform_tid，同时分销订单号的供应商必须和开票的授权账号一致
func (r *AlibabaEinvoiceCreatereqRequest) SetDistributionTid(_distributionTid string) error {
    r._distributionTid = _distributionTid
    r.Set("distribution_tid", _distributionTid)
    return nil
}

// DistributionTid Getter
func (r AlibabaEinvoiceCreatereqRequest) GetDistributionTid() string {
    return r._distributionTid
}
