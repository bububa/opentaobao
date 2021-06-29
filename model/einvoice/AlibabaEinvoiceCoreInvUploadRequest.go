package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-发票结果回传 APIRequest
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

func NewAlibabaEinvoiceCoreInvUploadRequest() *AlibabaEinvoiceCoreInvUploadRequest{
    return &AlibabaEinvoiceCoreInvUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetApiMethodName() string {
    return "alibaba.einvoice.core.inv.upload"
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSumTax(sumTax string) error {
    r.sumTax = sumTax
    r.Set("sum_tax", sumTax)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetSumTax() string {
    return r.sumTax
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeePhone(payeePhone string) error {
    r.payeePhone = payeePhone
    r.Set("payee_phone", payeePhone)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeePhone() string {
    return r.payeePhone
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformTid() string {
    return r.platformTid
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetSumPrice() string {
    return r.sumPrice
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceAmount(invoiceAmount string) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceAmount() string {
    return r.invoiceAmount
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetSource() string {
    return r.source
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceFileType(invoiceFileType string) error {
    r.invoiceFileType = invoiceFileType
    r.Set("invoice_file_type", invoiceFileType)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceFileType() string {
    return r.invoiceFileType
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeReceiver(payeeReceiver string) error {
    r.payeeReceiver = payeeReceiver
    r.Set("payee_receiver", payeeReceiver)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeReceiver() string {
    return r.payeeReceiver
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetRedNoticeNo(redNoticeNo string) error {
    r.redNoticeNo = redNoticeNo
    r.Set("red_notice_no", redNoticeNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetRedNoticeNo() string {
    return r.redNoticeNo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetReceiveMobile(receiveMobile string) error {
    r.receiveMobile = receiveMobile
    r.Set("receive_mobile", receiveMobile)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetReceiveMobile() string {
    return r.receiveMobile
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetApplyId() string {
    return r.applyId
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeName() string {
    return r.payeeName
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetQrCode(qrCode string) error {
    r.qrCode = qrCode
    r.Set("qr_code", qrCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetQrCode() string {
    return r.qrCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetLevyType(levyType string) error {
    r.levyType = levyType
    r.Set("levy_type", levyType)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetLevyType() string {
    return r.levyType
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceType(invoiceType string) error {
    r.invoiceType = invoiceType
    r.Set("invoice_type", invoiceType)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceType() string {
    return r.invoiceType
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerName(payerName string) error {
    r.payerName = payerName
    r.Set("payer_name", payerName)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerName() string {
    return r.payerName
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceMemo(invoiceMemo string) error {
    r.invoiceMemo = invoiceMemo
    r.Set("invoice_memo", invoiceMemo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceMemo() string {
    return r.invoiceMemo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerEmail(payerEmail string) error {
    r.payerEmail = payerEmail
    r.Set("payer_email", payerEmail)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerEmail() string {
    return r.payerEmail
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetAntiFakeCode(antiFakeCode string) error {
    r.antiFakeCode = antiFakeCode
    r.Set("anti_fake_code", antiFakeCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetAntiFakeCode() string {
    return r.antiFakeCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerBankAccountId(payerBankAccountId string) error {
    r.payerBankAccountId = payerBankAccountId
    r.Set("payer_bank_account_id", payerBankAccountId)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerBankAccountId() string {
    return r.payerBankAccountId
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerBankName(payerBankName string) error {
    r.payerBankName = payerBankName
    r.Set("payer_bank_name", payerBankName)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerBankName() string {
    return r.payerBankName
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeBankAccountId(payeeBankAccountId string) error {
    r.payeeBankAccountId = payeeBankAccountId
    r.Set("payee_bank_account_id", payeeBankAccountId)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeBankAccountId() string {
    return r.payeeBankAccountId
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeChecker(payeeChecker string) error {
    r.payeeChecker = payeeChecker
    r.Set("payee_checker", payeeChecker)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeChecker() string {
    return r.payeeChecker
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerPhone(payerPhone string) error {
    r.payerPhone = payerPhone
    r.Set("payer_phone", payerPhone)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerPhone() string {
    return r.payerPhone
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetNormalInvoiceCode(normalInvoiceCode string) error {
    r.normalInvoiceCode = normalInvoiceCode
    r.Set("normal_invoice_code", normalInvoiceCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetNormalInvoiceCode() string {
    return r.normalInvoiceCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetDeviceNo() string {
    return r.deviceNo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceDate() string {
    return r.invoiceDate
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetCheckCode(checkCode string) error {
    r.checkCode = checkCode
    r.Set("check_code", checkCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetCheckCode() string {
    return r.checkCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeAddress(payeeAddress string) error {
    r.payeeAddress = payeeAddress
    r.Set("payee_address", payeeAddress)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeAddress() string {
    return r.payeeAddress
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetNormalInvoiceNo(normalInvoiceNo string) error {
    r.normalInvoiceNo = normalInvoiceNo
    r.Set("normal_invoice_no", normalInvoiceNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetNormalInvoiceNo() string {
    return r.normalInvoiceNo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerAddress(payerAddress string) error {
    r.payerAddress = payerAddress
    r.Set("payer_address", payerAddress)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerAddress() string {
    return r.payerAddress
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeBankName(payeeBankName string) error {
    r.payeeBankName = payeeBankName
    r.Set("payee_bank_name", payeeBankName)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeBankName() string {
    return r.payeeBankName
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceItems(invoiceItems []InvoiceResultItemDto) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceItems() []InvoiceResultItemDto {
    return r.invoiceItems
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayeeOperator(payeeOperator string) error {
    r.payeeOperator = payeeOperator
    r.Set("payee_operator", payeeOperator)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayeeOperator() string {
    return r.payeeOperator
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceFileData(invoiceFileData []*model.File) error {
    r.invoiceFileData = invoiceFileData
    r.Set("invoice_file_data", invoiceFileData)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceFileData() []*model.File {
    return r.invoiceFileData
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPayerRegisterNo(payerRegisterNo string) error {
    r.payerRegisterNo = payerRegisterNo
    r.Set("payer_register_no", payerRegisterNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPayerRegisterNo() string {
    return r.payerRegisterNo
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBusinessType(businessType int64) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetBusinessType() int64 {
    return r.businessType
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSpecialFlag(specialFlag string) error {
    r.specialFlag = specialFlag
    r.Set("special_flag", specialFlag)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetSpecialFlag() string {
    return r.specialFlag
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformCode() string {
    return r.platformCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetPlatformUserId() string {
    return r.platformUserId
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBizErrorMsg(bizErrorMsg string) error {
    r.bizErrorMsg = bizErrorMsg
    r.Set("biz_error_msg", bizErrorMsg)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetBizErrorMsg() string {
    return r.bizErrorMsg
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetBizErrorCode(bizErrorCode string) error {
    r.bizErrorCode = bizErrorCode
    r.Set("biz_error_code", bizErrorCode)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetBizErrorCode() string {
    return r.bizErrorCode
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetCreateResult(createResult string) error {
    r.createResult = createResult
    r.Set("create_result", createResult)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetCreateResult() string {
    return r.createResult
}

func (r *AlibabaEinvoiceCoreInvUploadRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaEinvoiceCoreInvUploadRequest) GetSerialNo() string {
    return r.serialNo
}

