package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传发票查验的结果 APIRequest
alibaba.einvoice.income.verify.return

服务商回传发票查验的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的查验回传
*/
type AlibabaEinvoiceIncomeVerifyReturnRequest struct {
    model.Params

    // 校验码，success=true时必填
    checksum   string 

    // 错误码，success=false时必填
    errorCode   string 

    // 错误信息，success=false时必填
    errorMessage   string 

    // 发票影像编号，type=1时必填
    imageId   string 

    // 价税合计金额，success=true时必填，invoiceAmount=sumPrice+sumTax
    invoiceAmount   string 

    // 发票代码，success=true时必填
    invoiceCode   string 

    // 开票日期，格式为yyyy-MM-dd，success=true时必填
    invoiceDate   string 

    // 发票明细
    invoiceItems   []InvoiceItem 

    // 发票备注
    invoiceMemo   string 

    // 发票号码，success=true时必填
    invoiceNo   string 

    // 机器编号
    machineNo   string 

    // 销售方地址电话
    payeeAddressPhone   string 

    // 销售方银行及账号
    payeeBankAccount   string 

    // 复核人
    payeeChecker   string 

    // 销售方名称，success=true时必填
    payeeName   string 

    // 开票人
    payeeOperator   string 

    // 收款人
    payeeReceiver   string 

    // 销售方纳税人识别号，success=true时必填
    payeeRegisterNo   string 

    // 购买方地址电话
    payerAddressPhone   string 

    // 购买方银行及账号
    payerBankAccount   string 

    // 购买方名称，即发票抬头，success=true时必填
    payerName   string 

    // 购买方纳税人识别号
    payerRegisterNo   string 

    // 开票请求标识，扫描驱动回传type=1时填批次号
    reqIndex   string 

    // 查验结果，true=成功，false=失败
    success   bool 

    // 发票不含税金额，success=true时必填
    sumPrice   string 

    // 发票税额，success=true时必填
    sumTax   string 

    // 请求类型，0=阿里主动发起的查验，1=扫描驱动服务商主动回传查验结果
    type   int64 

    // 发票状态，0=无效（作废），1=有效
    invoiceStatus   int64 

}

func NewAlibabaEinvoiceIncomeVerifyReturnRequest() *AlibabaEinvoiceIncomeVerifyReturnRequest{
    return &AlibabaEinvoiceIncomeVerifyReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.verify.return"
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetChecksum(checksum string) error {
    r.checksum = checksum
    r.Set("checksum", checksum)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetChecksum() string {
    return r.checksum
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetErrorCode() string {
    return r.errorCode
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetImageId(imageId string) error {
    r.imageId = imageId
    r.Set("image_id", imageId)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetImageId() string {
    return r.imageId
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceAmount(invoiceAmount string) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceAmount() string {
    return r.invoiceAmount
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceDate() string {
    return r.invoiceDate
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceItems(invoiceItems []InvoiceItem) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceItems() []InvoiceItem {
    return r.invoiceItems
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceMemo(invoiceMemo string) error {
    r.invoiceMemo = invoiceMemo
    r.Set("invoice_memo", invoiceMemo)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceMemo() string {
    return r.invoiceMemo
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetMachineNo(machineNo string) error {
    r.machineNo = machineNo
    r.Set("machine_no", machineNo)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetMachineNo() string {
    return r.machineNo
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeAddressPhone(payeeAddressPhone string) error {
    r.payeeAddressPhone = payeeAddressPhone
    r.Set("payee_address_phone", payeeAddressPhone)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeAddressPhone() string {
    return r.payeeAddressPhone
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeBankAccount(payeeBankAccount string) error {
    r.payeeBankAccount = payeeBankAccount
    r.Set("payee_bank_account", payeeBankAccount)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeBankAccount() string {
    return r.payeeBankAccount
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeChecker(payeeChecker string) error {
    r.payeeChecker = payeeChecker
    r.Set("payee_checker", payeeChecker)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeChecker() string {
    return r.payeeChecker
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeName() string {
    return r.payeeName
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeOperator(payeeOperator string) error {
    r.payeeOperator = payeeOperator
    r.Set("payee_operator", payeeOperator)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeOperator() string {
    return r.payeeOperator
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeReceiver(payeeReceiver string) error {
    r.payeeReceiver = payeeReceiver
    r.Set("payee_receiver", payeeReceiver)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeReceiver() string {
    return r.payeeReceiver
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayerAddressPhone(payerAddressPhone string) error {
    r.payerAddressPhone = payerAddressPhone
    r.Set("payer_address_phone", payerAddressPhone)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayerAddressPhone() string {
    return r.payerAddressPhone
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayerBankAccount(payerBankAccount string) error {
    r.payerBankAccount = payerBankAccount
    r.Set("payer_bank_account", payerBankAccount)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayerBankAccount() string {
    return r.payerBankAccount
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayerName(payerName string) error {
    r.payerName = payerName
    r.Set("payer_name", payerName)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayerName() string {
    return r.payerName
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetPayerRegisterNo(payerRegisterNo string) error {
    r.payerRegisterNo = payerRegisterNo
    r.Set("payer_register_no", payerRegisterNo)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetPayerRegisterNo() string {
    return r.payerRegisterNo
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetReqIndex() string {
    return r.reqIndex
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetSuccess() bool {
    return r.success
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetSumPrice() string {
    return r.sumPrice
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetSumTax(sumTax string) error {
    r.sumTax = sumTax
    r.Set("sum_tax", sumTax)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetSumTax() string {
    return r.sumTax
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaEinvoiceIncomeVerifyReturnRequest) SetInvoiceStatus(invoiceStatus int64) error {
    r.invoiceStatus = invoiceStatus
    r.Set("invoice_status", invoiceStatus)
    return nil
}

func (r AlibabaEinvoiceIncomeVerifyReturnRequest) GetInvoiceStatus() int64 {
    return r.invoiceStatus
}

