package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸票打印接口 API请求
alibaba.einvoice.paper.print

打印一张已开具成功的纸票
*/
type AlibabaEinvoicePaperPrintRequest struct {
    model.Params
    // 打印框设置，0=不弹打印设置框，1=弹出打印设置框
    dialogSettingFlag   int64
    // 是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
    forcePrint   bool
    // 销售方纳税人识别号
    payeeRegisterNo   string
    // 打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
    printFlag   int64
    // 开票流水号
    serialNo   string
}

// 初始化AlibabaEinvoicePaperPrintRequest对象
func NewAlibabaEinvoicePaperPrintRequest() *AlibabaEinvoicePaperPrintRequest{
    return &AlibabaEinvoicePaperPrintRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperPrintRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.print"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperPrintRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DialogSettingFlag Setter
// 打印框设置，0=不弹打印设置框，1=弹出打印设置框
func (r *AlibabaEinvoicePaperPrintRequest) SetDialogSettingFlag(dialogSettingFlag int64) error {
    r.dialogSettingFlag = dialogSettingFlag
    r.Set("dialog_setting_flag", dialogSettingFlag)
    return nil
}

// DialogSettingFlag Getter
func (r AlibabaEinvoicePaperPrintRequest) GetDialogSettingFlag() int64 {
    return r.dialogSettingFlag
}
// ForcePrint Setter
// 是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
func (r *AlibabaEinvoicePaperPrintRequest) SetForcePrint(forcePrint bool) error {
    r.forcePrint = forcePrint
    r.Set("force_print", forcePrint)
    return nil
}

// ForcePrint Getter
func (r AlibabaEinvoicePaperPrintRequest) GetForcePrint() bool {
    return r.forcePrint
}
// PayeeRegisterNo Setter
// 销售方纳税人识别号
func (r *AlibabaEinvoicePaperPrintRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoicePaperPrintRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// PrintFlag Setter
// 打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
func (r *AlibabaEinvoicePaperPrintRequest) SetPrintFlag(printFlag int64) error {
    r.printFlag = printFlag
    r.Set("print_flag", printFlag)
    return nil
}

// PrintFlag Getter
func (r AlibabaEinvoicePaperPrintRequest) GetPrintFlag() int64 {
    return r.printFlag
}
// SerialNo Setter
// 开票流水号
func (r *AlibabaEinvoicePaperPrintRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoicePaperPrintRequest) GetSerialNo() string {
    return r.serialNo
}
