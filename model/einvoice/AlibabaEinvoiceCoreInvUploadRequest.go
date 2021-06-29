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
    _sumTax   string
    // 销方联系电话。
    _payeePhone   string
    // 业务平台发票申请对应的订单号。  用于source=upload时区分业务平台订单号。  source=apply时可空
    _platformTid   string
    // 合计金额（不含税），格式为2位小数。  当开红票时，该字段为负数
    _sumPrice   string
    // 合计含税金额（开票金额），格式为2位小数。  当开红票时，该字段为负数。
    _invoiceAmount   string
    // 发票来源，可选值：  apply: 间连模式，服务商基于申请开具的发票；async：直连开票模式，ISV回传开票结果；upload：直接回传，进行归集与交付的发票；
    _source   string
    // 发票板式文件类型。可选值：  PDF, OFD。  电票时必传。
    _invoiceFileType   string
    // 收款人
    _payeeReceiver   string
    // 红字通知单号
    _redNoticeNo   string
    // 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。
    _payeeRegisterNo   string
    // 购方手机号码，用于收票
    _receiveMobile   string
    // 发票申请ID, 由阿里发票平台生成。  source=apply时 必填。
    _applyId   string
    // 销方名称
    _payeeName   string
    // 二维码
    _qrCode   string
    // 征税方式，0普通征收，1减按征收，2差额征收
    _levyType   string
    // 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
    _invoiceType   string
    // 购方抬头
    _payerName   string
    // 发票号码
    _invoiceNo   string
    // 发票备注，会显示在票面
    _invoiceMemo   string
    // 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
    _payerEmail   string
    // 发票防伪码/密码
    _antiFakeCode   string
    // 购方银行账号， 专票必填。
    _payerBankAccountId   string
    // 购方开户行名称，  专票必填。
    _payerBankName   string
    // 销方银行账号
    _payeeBankAccountId   string
    // 复核人
    _payeeChecker   string
    // 购方联系电话，  专票必填。
    _payerPhone   string
    // 原发票代码(开红票时必须)
    _normalInvoiceCode   string
    // 开票分机号/机器编号
    _deviceNo   string
    // 开票日期，格式 yyyy-MM-dd
    _invoiceDate   string
    // 发票代码
    _invoiceCode   string
    // 校验码
    _checkCode   string
    // 销方地址。
    _payeeAddress   string
    // 原发票号码(开红票时必须)
    _normalInvoiceNo   string
    // 购方地址，  专票必填。
    _payerAddress   string
    // 销方开户行名称。
    _payeeBankName   string
    // 发票明细。source=apply时必填，其他source可为空
    _invoiceItems   []InvoiceResultItemDTO
    // 开票人
    _payeeOperator   string
    // 发票板式文件数据，字节数据。  电票时必传。
    _invoiceFileData   *model.File
    // 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
    _payerRegisterNo   string
    // 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
    _invoiceKind   int64
    // 抬头类型。可选值：  0：个人  1：企业
    _businessType   int64
    // 特殊票种标识，可选值：  02: 农产品收购票
    _specialFlag   string
    // 业务平台Code, 由发票中台分配。  用于source=upload时标识需交付发票的业务平台。  source=apply时可空
    _platformCode   string
    // 业务平台uid
    _platformUserId   string
    // 开票失败错误信息，  开票失败(create_result=fail)时必填。
    _bizErrorMsg   string
    // 开票失败错误码，  开票失败(create_result=fail)时必填。
    _bizErrorCode   string
    // 开票结果，枚举值：  success: 发票开具成功；  fail: 开票失败；  source=async时必填，传实际的开票结果。其他source可不传，默认为success
    _createResult   string
    // 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。  source=async时必填，其他source可为空
    _serialNo   string
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
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSumTax(_sumTax string) error {
    r._sumTax = _sumTax
    r.Set("sum_tax", _sumTax)
    return nil
}

// SumTax Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSumTax() string {
    return r._sumTax
}
// PayeePhone Setter
// 销方联系电话。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeePhone(_payeePhone string) error {
    r._payeePhone = _payeePhone
    r.Set("payee_phone", _payeePhone)
    return nil
}

// PayeePhone Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeePhone() string {
    return r._payeePhone
}
// PlatformTid Setter
// 业务平台发票申请对应的订单号。  用于source=upload时区分业务平台订单号。  source=apply时可空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformTid(_platformTid string) error {
    r._platformTid = _platformTid
    r.Set("platform_tid", _platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformTid() string {
    return r._platformTid
}
// SumPrice Setter
// 合计金额（不含税），格式为2位小数。  当开红票时，该字段为负数
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSumPrice(_sumPrice string) error {
    r._sumPrice = _sumPrice
    r.Set("sum_price", _sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSumPrice() string {
    return r._sumPrice
}
// InvoiceAmount Setter
// 合计含税金额（开票金额），格式为2位小数。  当开红票时，该字段为负数。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceAmount(_invoiceAmount string) error {
    r._invoiceAmount = _invoiceAmount
    r.Set("invoice_amount", _invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceAmount() string {
    return r._invoiceAmount
}
// Source Setter
// 发票来源，可选值：  apply: 间连模式，服务商基于申请开具的发票；async：直连开票模式，ISV回传开票结果；upload：直接回传，进行归集与交付的发票；
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSource() string {
    return r._source
}
// InvoiceFileType Setter
// 发票板式文件类型。可选值：  PDF, OFD。  电票时必传。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceFileType(_invoiceFileType string) error {
    r._invoiceFileType = _invoiceFileType
    r.Set("invoice_file_type", _invoiceFileType)
    return nil
}

// InvoiceFileType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceFileType() string {
    return r._invoiceFileType
}
// PayeeReceiver Setter
// 收款人
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeReceiver(_payeeReceiver string) error {
    r._payeeReceiver = _payeeReceiver
    r.Set("payee_receiver", _payeeReceiver)
    return nil
}

// PayeeReceiver Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeReceiver() string {
    return r._payeeReceiver
}
// RedNoticeNo Setter
// 红字通知单号
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetRedNoticeNo(_redNoticeNo string) error {
    r._redNoticeNo = _redNoticeNo
    r.Set("red_notice_no", _redNoticeNo)
    return nil
}

// RedNoticeNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetRedNoticeNo() string {
    return r._redNoticeNo
}
// PayeeRegisterNo Setter
// 销方税务登记证号。由大写字母或数字组成，长度要求15~20位。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// ReceiveMobile Setter
// 购方手机号码，用于收票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetReceiveMobile(_receiveMobile string) error {
    r._receiveMobile = _receiveMobile
    r.Set("receive_mobile", _receiveMobile)
    return nil
}

// ReceiveMobile Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetReceiveMobile() string {
    return r._receiveMobile
}
// ApplyId Setter
// 发票申请ID, 由阿里发票平台生成。  source=apply时 必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetApplyId() string {
    return r._applyId
}
// PayeeName Setter
// 销方名称
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeName(_payeeName string) error {
    r._payeeName = _payeeName
    r.Set("payee_name", _payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeName() string {
    return r._payeeName
}
// QrCode Setter
// 二维码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetQrCode(_qrCode string) error {
    r._qrCode = _qrCode
    r.Set("qr_code", _qrCode)
    return nil
}

// QrCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetQrCode() string {
    return r._qrCode
}
// LevyType Setter
// 征税方式，0普通征收，1减按征收，2差额征收
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetLevyType(_levyType string) error {
    r._levyType = _levyType
    r.Set("levy_type", _levyType)
    return nil
}

// LevyType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetLevyType() string {
    return r._levyType
}
// InvoiceType Setter
// 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceType(_invoiceType string) error {
    r._invoiceType = _invoiceType
    r.Set("invoice_type", _invoiceType)
    return nil
}

// InvoiceType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceType() string {
    return r._invoiceType
}
// PayerName Setter
// 购方抬头
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerName(_payerName string) error {
    r._payerName = _payerName
    r.Set("payer_name", _payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerName() string {
    return r._payerName
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// InvoiceMemo Setter
// 发票备注，会显示在票面
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceMemo(_invoiceMemo string) error {
    r._invoiceMemo = _invoiceMemo
    r.Set("invoice_memo", _invoiceMemo)
    return nil
}

// InvoiceMemo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceMemo() string {
    return r._invoiceMemo
}
// PayerEmail Setter
// 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerEmail(_payerEmail string) error {
    r._payerEmail = _payerEmail
    r.Set("payer_email", _payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerEmail() string {
    return r._payerEmail
}
// AntiFakeCode Setter
// 发票防伪码/密码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetAntiFakeCode(_antiFakeCode string) error {
    r._antiFakeCode = _antiFakeCode
    r.Set("anti_fake_code", _antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetAntiFakeCode() string {
    return r._antiFakeCode
}
// PayerBankAccountId Setter
// 购方银行账号， 专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerBankAccountId(_payerBankAccountId string) error {
    r._payerBankAccountId = _payerBankAccountId
    r.Set("payer_bank_account_id", _payerBankAccountId)
    return nil
}

// PayerBankAccountId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerBankAccountId() string {
    return r._payerBankAccountId
}
// PayerBankName Setter
// 购方开户行名称，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerBankName(_payerBankName string) error {
    r._payerBankName = _payerBankName
    r.Set("payer_bank_name", _payerBankName)
    return nil
}

// PayerBankName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerBankName() string {
    return r._payerBankName
}
// PayeeBankAccountId Setter
// 销方银行账号
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeBankAccountId(_payeeBankAccountId string) error {
    r._payeeBankAccountId = _payeeBankAccountId
    r.Set("payee_bank_account_id", _payeeBankAccountId)
    return nil
}

// PayeeBankAccountId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeBankAccountId() string {
    return r._payeeBankAccountId
}
// PayeeChecker Setter
// 复核人
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeChecker(_payeeChecker string) error {
    r._payeeChecker = _payeeChecker
    r.Set("payee_checker", _payeeChecker)
    return nil
}

// PayeeChecker Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeChecker() string {
    return r._payeeChecker
}
// PayerPhone Setter
// 购方联系电话，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerPhone(_payerPhone string) error {
    r._payerPhone = _payerPhone
    r.Set("payer_phone", _payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerPhone() string {
    return r._payerPhone
}
// NormalInvoiceCode Setter
// 原发票代码(开红票时必须)
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
    r._normalInvoiceCode = _normalInvoiceCode
    r.Set("normal_invoice_code", _normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetNormalInvoiceCode() string {
    return r._normalInvoiceCode
}
// DeviceNo Setter
// 开票分机号/机器编号
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetDeviceNo(_deviceNo string) error {
    r._deviceNo = _deviceNo
    r.Set("device_no", _deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetDeviceNo() string {
    return r._deviceNo
}
// InvoiceDate Setter
// 开票日期，格式 yyyy-MM-dd
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceDate(_invoiceDate string) error {
    r._invoiceDate = _invoiceDate
    r.Set("invoice_date", _invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceDate() string {
    return r._invoiceDate
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// CheckCode Setter
// 校验码
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetCheckCode(_checkCode string) error {
    r._checkCode = _checkCode
    r.Set("check_code", _checkCode)
    return nil
}

// CheckCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetCheckCode() string {
    return r._checkCode
}
// PayeeAddress Setter
// 销方地址。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeAddress(_payeeAddress string) error {
    r._payeeAddress = _payeeAddress
    r.Set("payee_address", _payeeAddress)
    return nil
}

// PayeeAddress Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeAddress() string {
    return r._payeeAddress
}
// NormalInvoiceNo Setter
// 原发票号码(开红票时必须)
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
    r._normalInvoiceNo = _normalInvoiceNo
    r.Set("normal_invoice_no", _normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetNormalInvoiceNo() string {
    return r._normalInvoiceNo
}
// PayerAddress Setter
// 购方地址，  专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerAddress(_payerAddress string) error {
    r._payerAddress = _payerAddress
    r.Set("payer_address", _payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerAddress() string {
    return r._payerAddress
}
// PayeeBankName Setter
// 销方开户行名称。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeBankName(_payeeBankName string) error {
    r._payeeBankName = _payeeBankName
    r.Set("payee_bank_name", _payeeBankName)
    return nil
}

// PayeeBankName Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeBankName() string {
    return r._payeeBankName
}
// InvoiceItems Setter
// 发票明细。source=apply时必填，其他source可为空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceItems(_invoiceItems []InvoiceResultItemDTO) error {
    r._invoiceItems = _invoiceItems
    r.Set("invoice_items", _invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceItems() []InvoiceResultItemDTO {
    return r._invoiceItems
}
// PayeeOperator Setter
// 开票人
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeOperator(_payeeOperator string) error {
    r._payeeOperator = _payeeOperator
    r.Set("payee_operator", _payeeOperator)
    return nil
}

// PayeeOperator Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeOperator() string {
    return r._payeeOperator
}
// InvoiceFileData Setter
// 发票板式文件数据，字节数据。  电票时必传。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceFileData(_invoiceFileData *model.File) error {
    r._invoiceFileData = _invoiceFileData
    r.Set("invoice_file_data", _invoiceFileData)
    return nil
}

// InvoiceFileData Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceFileData() *model.File {
    return r._invoiceFileData
}
// PayerRegisterNo Setter
// 购方税务登记证号，由大写字母或数字组成，长度要求15~20位。  开企业抬头时必填， 专票必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
    r._payerRegisterNo = _payerRegisterNo
    r.Set("payer_register_no", _payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerRegisterNo() string {
    return r._payerRegisterNo
}
// InvoiceKind Setter
// 开票发票类型  可选值：  0: 电票  1：纸质普票  2：纸质专票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceKind(_invoiceKind int64) error {
    r._invoiceKind = _invoiceKind
    r.Set("invoice_kind", _invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceKind() int64 {
    return r._invoiceKind
}
// BusinessType Setter
// 抬头类型。可选值：  0：个人  1：企业
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBusinessType(_businessType int64) error {
    r._businessType = _businessType
    r.Set("business_type", _businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetBusinessType() int64 {
    return r._businessType
}
// SpecialFlag Setter
// 特殊票种标识，可选值：  02: 农产品收购票
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSpecialFlag(_specialFlag string) error {
    r._specialFlag = _specialFlag
    r.Set("special_flag", _specialFlag)
    return nil
}

// SpecialFlag Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSpecialFlag() string {
    return r._specialFlag
}
// PlatformCode Setter
// 业务平台Code, 由发票中台分配。  用于source=upload时标识需交付发票的业务平台。  source=apply时可空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformCode() string {
    return r._platformCode
}
// PlatformUserId Setter
// 业务平台uid
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformUserId(_platformUserId string) error {
    r._platformUserId = _platformUserId
    r.Set("platform_user_id", _platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformUserId() string {
    return r._platformUserId
}
// BizErrorMsg Setter
// 开票失败错误信息，  开票失败(create_result=fail)时必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBizErrorMsg(_bizErrorMsg string) error {
    r._bizErrorMsg = _bizErrorMsg
    r.Set("biz_error_msg", _bizErrorMsg)
    return nil
}

// BizErrorMsg Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetBizErrorMsg() string {
    return r._bizErrorMsg
}
// BizErrorCode Setter
// 开票失败错误码，  开票失败(create_result=fail)时必填。
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBizErrorCode(_bizErrorCode string) error {
    r._bizErrorCode = _bizErrorCode
    r.Set("biz_error_code", _bizErrorCode)
    return nil
}

// BizErrorCode Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetBizErrorCode() string {
    return r._bizErrorCode
}
// CreateResult Setter
// 开票结果，枚举值：  success: 发票开具成功；  fail: 开票失败；  source=async时必填，传实际的开票结果。其他source可不传，默认为success
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetCreateResult(_createResult string) error {
    r._createResult = _createResult
    r.Set("create_result", _createResult)
    return nil
}

// CreateResult Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetCreateResult() string {
    return r._createResult
}
// SerialNo Setter
// 开票流水号/序列号，唯一标志一笔开票请求，由于阿里发票中台生成。  source=async时必填，其他source可为空
func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceCoreInvUploadRequest) GetSerialNo() string {
    return r._serialNo
}
