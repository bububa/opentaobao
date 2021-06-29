package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票商回传开票结果 API请求
alibaba.einvoice.partner.return

开票商返回开票结果数据
*/
type AlibabaEinvoicePartnerReturnRequest struct {
    model.Params
    // 电商平台身份标识码，TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
    _platformCode   string
    // 开票金额
    _invoiceAmount   string
    // 发票密文，密码区的字符串
    _ciphertext   string
    // 收款方税务登记证号
    _payeeRegisterNo   string
    // 二维码
    _qrCode   string
    // erp自定义单据号
    _erpTid   string
    // 文件类型(pdf,jpg,png)
    _fileDataType   string
    // 发票号码
    _invoiceNo   string
    // 发票日期
    _invoiceDate   string
    // 发票文件PDF内容，PDF的byte[]字段串。
    _invoiceFileData   *model.File
    // 流水号
    _serialNo   string
    // 防伪码
    _antiFakeCode   string
    // 税控设备编号(新版电子发票有)
    _deviceNo   string
    // 发票代码
    _invoiceCode   string
    // 电商平台对应的订单号
    _platformTid   string
    // 开票结果"success"或者"fail"
    _createResult   string
    // 错误码
    _bizErrorCode   string
    // 错误信息
    _bizErrorMsg   string
    // 开票请求的唯一索引
    _reqIndex   string
    // 开票时间，格式为HH:mm:ss
    _invoiceTime   string
}

// 初始化AlibabaEinvoicePartnerReturnRequest对象
func NewAlibabaEinvoicePartnerReturnRequest() *AlibabaEinvoicePartnerReturnRequest{
    return &AlibabaEinvoicePartnerReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePartnerReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.partner.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePartnerReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlatformCode Setter
// 电商平台身份标识码，TB=淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、OTHER=其他,  (只传英文编码)
func (r *AlibabaEinvoicePartnerReturnRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetPlatformCode() string {
    return r._platformCode
}
// InvoiceAmount Setter
// 开票金额
func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceAmount(_invoiceAmount string) error {
    r._invoiceAmount = _invoiceAmount
    r.Set("invoice_amount", _invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceAmount() string {
    return r._invoiceAmount
}
// Ciphertext Setter
// 发票密文，密码区的字符串
func (r *AlibabaEinvoicePartnerReturnRequest) SetCiphertext(_ciphertext string) error {
    r._ciphertext = _ciphertext
    r.Set("ciphertext", _ciphertext)
    return nil
}

// Ciphertext Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetCiphertext() string {
    return r._ciphertext
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoicePartnerReturnRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// QrCode Setter
// 二维码
func (r *AlibabaEinvoicePartnerReturnRequest) SetQrCode(_qrCode string) error {
    r._qrCode = _qrCode
    r.Set("qr_code", _qrCode)
    return nil
}

// QrCode Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetQrCode() string {
    return r._qrCode
}
// ErpTid Setter
// erp自定义单据号
func (r *AlibabaEinvoicePartnerReturnRequest) SetErpTid(_erpTid string) error {
    r._erpTid = _erpTid
    r.Set("erp_tid", _erpTid)
    return nil
}

// ErpTid Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetErpTid() string {
    return r._erpTid
}
// FileDataType Setter
// 文件类型(pdf,jpg,png)
func (r *AlibabaEinvoicePartnerReturnRequest) SetFileDataType(_fileDataType string) error {
    r._fileDataType = _fileDataType
    r.Set("file_data_type", _fileDataType)
    return nil
}

// FileDataType Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetFileDataType() string {
    return r._fileDataType
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// InvoiceDate Setter
// 发票日期
func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceDate(_invoiceDate string) error {
    r._invoiceDate = _invoiceDate
    r.Set("invoice_date", _invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceDate() string {
    return r._invoiceDate
}
// InvoiceFileData Setter
// 发票文件PDF内容，PDF的byte[]字段串。
func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceFileData(_invoiceFileData *model.File) error {
    r._invoiceFileData = _invoiceFileData
    r.Set("invoice_file_data", _invoiceFileData)
    return nil
}

// InvoiceFileData Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceFileData() *model.File {
    return r._invoiceFileData
}
// SerialNo Setter
// 流水号
func (r *AlibabaEinvoicePartnerReturnRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetSerialNo() string {
    return r._serialNo
}
// AntiFakeCode Setter
// 防伪码
func (r *AlibabaEinvoicePartnerReturnRequest) SetAntiFakeCode(_antiFakeCode string) error {
    r._antiFakeCode = _antiFakeCode
    r.Set("anti_fake_code", _antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetAntiFakeCode() string {
    return r._antiFakeCode
}
// DeviceNo Setter
// 税控设备编号(新版电子发票有)
func (r *AlibabaEinvoicePartnerReturnRequest) SetDeviceNo(_deviceNo string) error {
    r._deviceNo = _deviceNo
    r.Set("device_no", _deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetDeviceNo() string {
    return r._deviceNo
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// PlatformTid Setter
// 电商平台对应的订单号
func (r *AlibabaEinvoicePartnerReturnRequest) SetPlatformTid(_platformTid string) error {
    r._platformTid = _platformTid
    r.Set("platform_tid", _platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetPlatformTid() string {
    return r._platformTid
}
// CreateResult Setter
// 开票结果"success"或者"fail"
func (r *AlibabaEinvoicePartnerReturnRequest) SetCreateResult(_createResult string) error {
    r._createResult = _createResult
    r.Set("create_result", _createResult)
    return nil
}

// CreateResult Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetCreateResult() string {
    return r._createResult
}
// BizErrorCode Setter
// 错误码
func (r *AlibabaEinvoicePartnerReturnRequest) SetBizErrorCode(_bizErrorCode string) error {
    r._bizErrorCode = _bizErrorCode
    r.Set("biz_error_code", _bizErrorCode)
    return nil
}

// BizErrorCode Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetBizErrorCode() string {
    return r._bizErrorCode
}
// BizErrorMsg Setter
// 错误信息
func (r *AlibabaEinvoicePartnerReturnRequest) SetBizErrorMsg(_bizErrorMsg string) error {
    r._bizErrorMsg = _bizErrorMsg
    r.Set("biz_error_msg", _bizErrorMsg)
    return nil
}

// BizErrorMsg Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetBizErrorMsg() string {
    return r._bizErrorMsg
}
// ReqIndex Setter
// 开票请求的唯一索引
func (r *AlibabaEinvoicePartnerReturnRequest) SetReqIndex(_reqIndex string) error {
    r._reqIndex = _reqIndex
    r.Set("req_index", _reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetReqIndex() string {
    return r._reqIndex
}
// InvoiceTime Setter
// 开票时间，格式为HH:mm:ss
func (r *AlibabaEinvoicePartnerReturnRequest) SetInvoiceTime(_invoiceTime string) error {
    r._invoiceTime = _invoiceTime
    r.Set("invoice_time", _invoiceTime)
    return nil
}

// InvoiceTime Getter
func (r AlibabaEinvoicePartnerReturnRequest) GetInvoiceTime() string {
    return r._invoiceTime
}
