package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票商回传开票结果 APIRequest
alibaba.einvoice.partner.return

开票商返回开票结果数据
*/
type AlibabaEinvoicePartnerReturnRequest struct {
    model.Params

    // 电商平台身份标识码，TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    platformCode   string 

    // 开票金额
    invoiceAmount   string 

    // 发票密文，密码区的字符串
    ciphertext   string 

    // 收款方税务登记证号
    payeeRegisterNo   string 

    // 二维码
    qrCode   string 

    // erp自定义单据号
    erpTid   string 

    // 文件类型(pdf,jpg,png)
    fileDataType   string 

    // 发票号码
    invoiceNo   string 

    // 发票日期
    invoiceDate   string 

    // 发票文件PDF内容，PDF的byte[]字段串。
    invoiceFileData   []*model.File 

    // 流水号
    serialNo   string 

    // 防伪码
    antiFakeCode   string 

    // 税控设备编号(新版电子发票有)
    deviceNo   string 

    // 发票代码
    invoiceCode   string 

    // 电商平台对应的订单号
    platformTid   string 

    // 开票结果"success"或者"fail"
    createResult   string 

    // 错误码
    bizErrorCode   string 

    // 错误信息
    bizErrorMsg   string 

    // 开票请求的唯一索引
    reqIndex   string 

    // 开票时间，格式为HH:mm:ss
    invoiceTime   string 

}

func NewAlibabaEinvoicePartnerReturnRequest() *AlibabaEinvoicePartnerReturnRequest{
    return &AlibabaEinvoicePartnerReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoicePartnerReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.partner.return"
}

func (r AlibabaEinvoicePartnerReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoicePartnerReturnRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetPlatformCode() string {
    return r.platformCode
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceAmount(invoiceAmount string) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceAmount() string {
    return r.invoiceAmount
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetCiphertext(ciphertext string) error {
    r.ciphertext = ciphertext
    r.Set("ciphertext", ciphertext)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetCiphertext() string {
    return r.ciphertext
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetQrCode(qrCode string) error {
    r.qrCode = qrCode
    r.Set("qr_code", qrCode)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetQrCode() string {
    return r.qrCode
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetErpTid(erpTid string) error {
    r.erpTid = erpTid
    r.Set("erp_tid", erpTid)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetErpTid() string {
    return r.erpTid
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetFileDataType(fileDataType string) error {
    r.fileDataType = fileDataType
    r.Set("file_data_type", fileDataType)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetFileDataType() string {
    return r.fileDataType
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceDate() string {
    return r.invoiceDate
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceFileData(invoiceFileData []*model.File) error {
    r.invoiceFileData = invoiceFileData
    r.Set("invoice_file_data", invoiceFileData)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceFileData() []*model.File {
    return r.invoiceFileData
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetSerialNo() string {
    return r.serialNo
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetAntiFakeCode(antiFakeCode string) error {
    r.antiFakeCode = antiFakeCode
    r.Set("anti_fake_code", antiFakeCode)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetAntiFakeCode() string {
    return r.antiFakeCode
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetDeviceNo() string {
    return r.deviceNo
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetPlatformTid() string {
    return r.platformTid
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetCreateResult(createResult string) error {
    r.createResult = createResult
    r.Set("create_result", createResult)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetCreateResult() string {
    return r.createResult
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetBizErrorCode(bizErrorCode string) error {
    r.bizErrorCode = bizErrorCode
    r.Set("biz_error_code", bizErrorCode)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetBizErrorCode() string {
    return r.bizErrorCode
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetBizErrorMsg(bizErrorMsg string) error {
    r.bizErrorMsg = bizErrorMsg
    r.Set("biz_error_msg", bizErrorMsg)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetBizErrorMsg() string {
    return r.bizErrorMsg
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetReqIndex() string {
    return r.reqIndex
}

func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceTime(invoiceTime string) error {
    r.invoiceTime = invoiceTime
    r.Set("invoice_time", invoiceTime)
    return nil
}

func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceTime() string {
    return r.invoiceTime
}

