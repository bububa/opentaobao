package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-发票结果回传 API请求
alibaba.einvoice.core.inv.upload

发票回传接口适用于以下场景：
① 阿里发票平台向ISV提交原始发票申请，ISV开具发票成功后，基于申请ID(apply_id)回传发票至阿里发票平台进行归集与交付。
② 直接回传发票给阿里发票平台，进行归集，并交付给业务前台和用户。
*/
type AlibabaEinvoiceCoreInvUploadRequest struct {
    model.Params
    // 合计税额，格式为2位小数。  当开红票时，该字段为负数
    sumTax   string
    // 销方联系电话。
    payeePhone   string
    // 业务平台发票申请对应的订单号。  用于source=upload时区分业务平台订单号。  source=apply时可空
    platformTid   string
    // 合计金额（不含税），格式为2位小数。  当开红票时，该字段为负数
    sumPrice   string
    // 合计含税金额（开票金额），格式为2位小数。  当开红票时，该字段为负数。
    invoiceAmount   string
    // 发票来源，可选值：  apply: 间连模式，服务商基于申请开具的发票；async：直连开票模式，ISV回传开票结果；upload：直接回传，进行归集与交付的发票；
    source   string
    // 发票板式文件类型。可选值：  PDF, OFD。  电票时必传。
    invoiceFileType   string
    // 收款人
    payeeReceiver   string
    // 红字通知单号
    redNoticeNo   string
    // 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。
    payeeRegisterNo   string
    // 购方手机号码，用于收票
    receiveMobile   string
    // 发票申请ID, 由阿里发票平台生成。  source=apply时 必填。
    applyId   string
    // 销方名称
    payeeName   string
    // 二维码
    qrCode   string
    // 征税方式，0普通征收，1减按征收，2差额征收
    levyType   string
    // 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
    invoiceType   string
    // 购方抬头
    payerName   string
    // 发票号码
    invoiceNo   string
    // 发票备注，会显示在票面
    invoiceMemo   string
    // 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
    payerEmail   string
    // 发票防伪码/密码
    antiFakeCode   string
    // 购方银行账号， 专票必填。
    payerBankAccountId   string
    // 购方开户行名称，  专票必填。
    payerBankName   string
    // 销方银行账号
    payeeBankAccountId   string
    // 复核人
    payeeChecker   string
    // 购方联系电话，  专票必填。
    payerPhone   string
    // 原发票代码(开红票时必须)
    normalInvoiceCode   string
    // 开票分机号/机器编号
    deviceNo   string
    // 开票日期，格式 yyyy-MM-dd
    invoiceDate   string
    // 发票代码
    invoiceCode   string
    // 校验码
    checkCode   string
    // 销方地址。
    payeeAddress   string
    // 原发票号码(开红票时必须)
    normalInvoiceNo   string
    // 购方地址，  专票必填。
    payerAddress   string
    // 销方开户行名称。
    payeeBankName   string
    // 发票明细。source=apply时必填，其他source可为空
    invoiceItems   []InvoiceResultItemDto
    // 开票人
    payeeOperator   string
    // 发票板式文件数据，字节数据。  电票时必传。
    invoiceFileData   []*model.File
    // 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
    payerRegisterNo   string
    // 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
    invoiceKind   int64
    // 抬头类型。可选值：  0：个人  1：企业
    businessType   int64
    // 特殊票种标识，可选值：  02: 农产品收购票
    specialFlag   string
    // 业务平台Code, 由发票中台分配。  用于source=upload时标识需交付发票的业务平台。  source=apply时可空
    platformCode   string
    // 业务平台uid
    platformUserId   string
    // 开票失败错误信息，  开票失败(create_result=fail)时必填。
    bizErrorMsg   string
    // 开票失败错误码，  开票失败(create_result=fail)时必填。
    bizErrorCode   string
    // 开票结果，枚举值：  success: 发票开具成功；  fail: 开票失败；  source=async时必填，传实际的开票结果。其他source可不传，默认为success
    createResult   string
    // 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。  source=async时必填，其他source可为空
    serialNo   string
}

// 初始化AlibabaEinvoiceCoreInvUploadRequest对象
func NewAlibabaEinvoiceCoreInvUploadRequest() *AlibabaEinvoiceCoreInvUploadRequest{
    return &AlibabaEinvoiceCoreInvUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCoreInvUploadRequest) GetApiMethodName() string {
    return "alibaba.einvoice.core.inv.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCoreInvUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SumTax Setter
// 合计税额，格式为2位小数。  当开红票时，该字段为负数
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSumTax(sumTax string) error {
    r.sumTax = sumTax
    r.Set("sum_tax", sumTax)
    return nil
}

// SumTax Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSumTax() string {
    return r.sumTax
}
// PayeePhone Setter
// 销方联系电话。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeePhone(payeePhone string) error {
    r.payeePhone = payeePhone
    r.Set("payee_phone", payeePhone)
    return nil
}

// PayeePhone Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeePhone() string {
    return r.payeePhone
}
// PlatformTid Setter
// 业务平台发票申请对应的订单号。  用于source=upload时区分业务平台订单号。  source=apply时可空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformTid() string {
    return r.platformTid
}
// SumPrice Setter
// 合计金额（不含税），格式为2位小数。  当开红票时，该字段为负数
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSumPrice() string {
    return r.sumPrice
}
// InvoiceAmount Setter
// 合计含税金额（开票金额），格式为2位小数。  当开红票时，该字段为负数。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceAmount(invoiceAmount string) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceAmount() string {
    return r.invoiceAmount
}
// Source Setter
// 发票来源，可选值：  apply: 间连模式，服务商基于申请开具的发票；async：直连开票模式，ISV回传开票结果；upload：直接回传，进行归集与交付的发票；
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSource() string {
    return r.source
}
// InvoiceFileType Setter
// 发票板式文件类型。可选值：  PDF, OFD。  电票时必传。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceFileType(invoiceFileType string) error {
    r.invoiceFileType = invoiceFileType
    r.Set("invoice_file_type", invoiceFileType)
    return nil
}

// InvoiceFileType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceFileType() string {
    return r.invoiceFileType
}
// PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeReceiver(payeeReceiver string) error {
    r.payeeReceiver = payeeReceiver
    r.Set("payee_receiver", payeeReceiver)
    return nil
}

// PayeeReceiver Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeReceiver() string {
    return r.payeeReceiver
}
// RedNoticeNo Setter
// 红字通知单号
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetRedNoticeNo(redNoticeNo string) error {
    r.redNoticeNo = redNoticeNo
    r.Set("red_notice_no", redNoticeNo)
    return nil
}

// RedNoticeNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetRedNoticeNo() string {
    return r.redNoticeNo
}
// PayeeRegisterNo Setter
// 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// ReceiveMobile Setter
// 购方手机号码，用于收票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetReceiveMobile(receiveMobile string) error {
    r.receiveMobile = receiveMobile
    r.Set("receive_mobile", receiveMobile)
    return nil
}

// ReceiveMobile Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetReceiveMobile() string {
    return r.receiveMobile
}
// ApplyId Setter
// 发票申请ID, 由阿里发票平台生成。  source=apply时 必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetApplyId() string {
    return r.applyId
}
// PayeeName Setter
// 销方名称
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeName() string {
    return r.payeeName
}
// QrCode Setter
// 二维码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetQrCode(qrCode string) error {
    r.qrCode = qrCode
    r.Set("qr_code", qrCode)
    return nil
}

// QrCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetQrCode() string {
    return r.qrCode
}
// LevyType Setter
// 征税方式，0普通征收，1减按征收，2差额征收
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetLevyType(levyType string) error {
    r.levyType = levyType
    r.Set("levy_type", levyType)
    return nil
}

// LevyType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetLevyType() string {
    return r.levyType
}
// InvoiceType Setter
// 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceType(invoiceType string) error {
    r.invoiceType = invoiceType
    r.Set("invoice_type", invoiceType)
    return nil
}

// InvoiceType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceType() string {
    return r.invoiceType
}
// PayerName Setter
// 购方抬头
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerName(payerName string) error {
    r.payerName = payerName
    r.Set("payer_name", payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerName() string {
    return r.payerName
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceNo() string {
    return r.invoiceNo
}
// InvoiceMemo Setter
// 发票备注，会显示在票面
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceMemo(invoiceMemo string) error {
    r.invoiceMemo = invoiceMemo
    r.Set("invoice_memo", invoiceMemo)
    return nil
}

// InvoiceMemo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceMemo() string {
    return r.invoiceMemo
}
// PayerEmail Setter
// 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerEmail(payerEmail string) error {
    r.payerEmail = payerEmail
    r.Set("payer_email", payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerEmail() string {
    return r.payerEmail
}
// AntiFakeCode Setter
// 发票防伪码/密码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetAntiFakeCode(antiFakeCode string) error {
    r.antiFakeCode = antiFakeCode
    r.Set("anti_fake_code", antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetAntiFakeCode() string {
    return r.antiFakeCode
}
// PayerBankAccountId Setter
// 购方银行账号， 专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerBankAccountId(payerBankAccountId string) error {
    r.payerBankAccountId = payerBankAccountId
    r.Set("payer_bank_account_id", payerBankAccountId)
    return nil
}

// PayerBankAccountId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerBankAccountId() string {
    return r.payerBankAccountId
}
// PayerBankName Setter
// 购方开户行名称，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerBankName(payerBankName string) error {
    r.payerBankName = payerBankName
    r.Set("payer_bank_name", payerBankName)
    return nil
}

// PayerBankName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerBankName() string {
    return r.payerBankName
}
// PayeeBankAccountId Setter
// 销方银行账号
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeBankAccountId(payeeBankAccountId string) error {
    r.payeeBankAccountId = payeeBankAccountId
    r.Set("payee_bank_account_id", payeeBankAccountId)
    return nil
}

// PayeeBankAccountId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeBankAccountId() string {
    return r.payeeBankAccountId
}
// PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeChecker(payeeChecker string) error {
    r.payeeChecker = payeeChecker
    r.Set("payee_checker", payeeChecker)
    return nil
}

// PayeeChecker Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeChecker() string {
    return r.payeeChecker
}
// PayerPhone Setter
// 购方联系电话，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerPhone(payerPhone string) error {
    r.payerPhone = payerPhone
    r.Set("payer_phone", payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerPhone() string {
    return r.payerPhone
}
// NormalInvoiceCode Setter
// 原发票代码(开红票时必须)
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetNormalInvoiceCode(normalInvoiceCode string) error {
    r.normalInvoiceCode = normalInvoiceCode
    r.Set("normal_invoice_code", normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetNormalInvoiceCode() string {
    return r.normalInvoiceCode
}
// DeviceNo Setter
// 开票分机号/机器编号
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetDeviceNo() string {
    return r.deviceNo
}
// InvoiceDate Setter
// 开票日期，格式 yyyy-MM-dd
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceDate() string {
    return r.invoiceDate
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceCode() string {
    return r.invoiceCode
}
// CheckCode Setter
// 校验码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetCheckCode(checkCode string) error {
    r.checkCode = checkCode
    r.Set("check_code", checkCode)
    return nil
}

// CheckCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetCheckCode() string {
    return r.checkCode
}
// PayeeAddress Setter
// 销方地址。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeAddress(payeeAddress string) error {
    r.payeeAddress = payeeAddress
    r.Set("payee_address", payeeAddress)
    return nil
}

// PayeeAddress Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeAddress() string {
    return r.payeeAddress
}
// NormalInvoiceNo Setter
// 原发票号码(开红票时必须)
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetNormalInvoiceNo(normalInvoiceNo string) error {
    r.normalInvoiceNo = normalInvoiceNo
    r.Set("normal_invoice_no", normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetNormalInvoiceNo() string {
    return r.normalInvoiceNo
}
// PayerAddress Setter
// 购方地址，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerAddress(payerAddress string) error {
    r.payerAddress = payerAddress
    r.Set("payer_address", payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerAddress() string {
    return r.payerAddress
}
// PayeeBankName Setter
// 销方开户行名称。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeBankName(payeeBankName string) error {
    r.payeeBankName = payeeBankName
    r.Set("payee_bank_name", payeeBankName)
    return nil
}

// PayeeBankName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeBankName() string {
    return r.payeeBankName
}
// InvoiceItems Setter
// 发票明细。source=apply时必填，其他source可为空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceItems(invoiceItems []InvoiceResultItemDto) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceItems() []InvoiceResultItemDto {
    return r.invoiceItems
}
// PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeOperator(payeeOperator string) error {
    r.payeeOperator = payeeOperator
    r.Set("payee_operator", payeeOperator)
    return nil
}

// PayeeOperator Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeOperator() string {
    return r.payeeOperator
}
// InvoiceFileData Setter
// 发票板式文件数据，字节数据。  电票时必传。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceFileData(invoiceFileData []*model.File) error {
    r.invoiceFileData = invoiceFileData
    r.Set("invoice_file_data", invoiceFileData)
    return nil
}

// InvoiceFileData Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceFileData() []*model.File {
    return r.invoiceFileData
}
// PayerRegisterNo Setter
// 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerRegisterNo(payerRegisterNo string) error {
    r.payerRegisterNo = payerRegisterNo
    r.Set("payer_register_no", payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerRegisterNo() string {
    return r.payerRegisterNo
}
// InvoiceKind Setter
// 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}
// BusinessType Setter
// 抬头类型。可选值：  0：个人  1：企业
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBusinessType(businessType int64) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetBusinessType() int64 {
    return r.businessType
}
// SpecialFlag Setter
// 特殊票种标识，可选值：  02: 农产品收购票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSpecialFlag(specialFlag string) error {
    r.specialFlag = specialFlag
    r.Set("special_flag", specialFlag)
    return nil
}

// SpecialFlag Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSpecialFlag() string {
    return r.specialFlag
}
// PlatformCode Setter
// 业务平台Code, 由发票中台分配。  用于source=upload时标识需交付发票的业务平台。  source=apply时可空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformCode() string {
    return r.platformCode
}
// PlatformUserId Setter
// 业务平台uid
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformUserId() string {
    return r.platformUserId
}
// BizErrorMsg Setter
// 开票失败错误信息，  开票失败(create_result=fail)时必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBizErrorMsg(bizErrorMsg string) error {
    r.bizErrorMsg = bizErrorMsg
    r.Set("biz_error_msg", bizErrorMsg)
    return nil
}

// BizErrorMsg Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetBizErrorMsg() string {
    return r.bizErrorMsg
}
// BizErrorCode Setter
// 开票失败错误码，  开票失败(create_result=fail)时必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBizErrorCode(bizErrorCode string) error {
    r.bizErrorCode = bizErrorCode
    r.Set("biz_error_code", bizErrorCode)
    return nil
}

// BizErrorCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetBizErrorCode() string {
    return r.bizErrorCode
}
// CreateResult Setter
// 开票结果，枚举值：  success: 发票开具成功；  fail: 开票失败；  source=async时必填，传实际的开票结果。其他source可不传，默认为success
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetCreateResult(createResult string) error {
    r.createResult = createResult
    r.Set("create_result", createResult)
    return nil
}

// CreateResult Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetCreateResult() string {
    return r.createResult
}
// SerialNo Setter
// 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。  source=async时必填，其他source可为空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSerialNo() string {
    return r.serialNo
}
