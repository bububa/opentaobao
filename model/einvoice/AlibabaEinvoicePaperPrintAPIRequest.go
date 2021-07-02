package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperPrintAPIRequest 纸票打印接口 API请求
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintAPIRequest struct {
	model.Params
	// 打印框设置，0=不弹打印设置框，1=弹出打印设置框
	_dialogSettingFlag int64
	// 是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
	_forcePrint bool
	// 销售方纳税人识别号
	_payeeRegisterNo string
	// 打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
	_printFlag int64
	// 开票流水号
	_serialNo string
}

// NewAlibabaEinvoicePaperPrintRequest 初始化AlibabaEinvoicePaperPrintAPIRequest对象
func NewAlibabaEinvoicePaperPrintRequest() *AlibabaEinvoicePaperPrintAPIRequest {
	return &AlibabaEinvoicePaperPrintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperPrintAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperPrintAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDialogSettingFlag is DialogSettingFlag Setter
// 打印框设置，0=不弹打印设置框，1=弹出打印设置框
func (r *AlibabaEinvoicePaperPrintAPIRequest) SetDialogSettingFlag(_dialogSettingFlag int64) error {
	r._dialogSettingFlag = _dialogSettingFlag
	r.Set("dialog_setting_flag", _dialogSettingFlag)
	return nil
}

// GetDialogSettingFlag DialogSettingFlag Getter
func (r AlibabaEinvoicePaperPrintAPIRequest) GetDialogSettingFlag() int64 {
	return r._dialogSettingFlag
}

// SetForcePrint is ForcePrint Setter
// 是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
func (r *AlibabaEinvoicePaperPrintAPIRequest) SetForcePrint(_forcePrint bool) error {
	r._forcePrint = _forcePrint
	r.Set("force_print", _forcePrint)
	return nil
}

// GetForcePrint ForcePrint Getter
func (r AlibabaEinvoicePaperPrintAPIRequest) GetForcePrint() bool {
	return r._forcePrint
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号
func (r *AlibabaEinvoicePaperPrintAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoicePaperPrintAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetPrintFlag is PrintFlag Setter
// 打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
func (r *AlibabaEinvoicePaperPrintAPIRequest) SetPrintFlag(_printFlag int64) error {
	r._printFlag = _printFlag
	r.Set("print_flag", _printFlag)
	return nil
}

// GetPrintFlag PrintFlag Getter
func (r AlibabaEinvoicePaperPrintAPIRequest) GetPrintFlag() int64 {
	return r._printFlag
}

// SetSerialNo is SerialNo Setter
// 开票流水号
func (r *AlibabaEinvoicePaperPrintAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaEinvoicePaperPrintAPIRequest) GetSerialNo() string {
	return r._serialNo
}
