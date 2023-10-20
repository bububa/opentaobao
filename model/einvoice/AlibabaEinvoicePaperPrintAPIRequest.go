package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicepaperprintAPIRequest 纸票打印接口 API请求
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
type AlibabaeinvoicepaperprintAPIRequest struct {
	model.Params
	// 销售方纳税人识别号
	_payeeRegisterNo string
	// 开票流水号
	_serialNo string
	// 打印框设置，0=不弹打印设置框，1=弹出打印设置框
	_dialogSettingFlag int64
	// 打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
	_printFlag int64
	// 是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
	_forcePrint bool
}

// NewAlibabaeinvoicepaperprintRequest 初始化AlibabaeinvoicepaperprintAPIRequest对象
func NewAlibabaeinvoicepaperprintRequest() *AlibabaeinvoicepaperprintAPIRequest {
	return &AlibabaeinvoicepaperprintAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicepaperprintAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicepaperprintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicepaperprintAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号
func (r *AlibabaeinvoicepaperprintAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicepaperprintAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetSerialNo is SerialNo Setter
// 开票流水号
func (r *AlibabaeinvoicepaperprintAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaeinvoicepaperprintAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetDialogSettingFlag is DialogSettingFlag Setter
// 打印框设置，0=不弹打印设置框，1=弹出打印设置框
func (r *AlibabaeinvoicepaperprintAPIRequest) SetDialogSettingFlag(_dialogSettingFlag int64) error {
	r._dialogSettingFlag = _dialogSettingFlag
	r.Set("dialog_setting_flag", _dialogSettingFlag)
	return nil
}

// GetDialogSettingFlag DialogSettingFlag Getter
func (r AlibabaeinvoicepaperprintAPIRequest) GetDialogSettingFlag() int64 {
	return r._dialogSettingFlag
}

// SetPrintFlag is PrintFlag Setter
// 打印标记，0=打印发票；1=打印清单。发票明细超过8行时会生成清单页，需要打印清单。
func (r *AlibabaeinvoicepaperprintAPIRequest) SetPrintFlag(_printFlag int64) error {
	r._printFlag = _printFlag
	r.Set("print_flag", _printFlag)
	return nil
}

// GetPrintFlag PrintFlag Getter
func (r AlibabaeinvoicepaperprintAPIRequest) GetPrintFlag() int64 {
	return r._printFlag
}

// SetForcePrint is ForcePrint Setter
// 是否强制打印，一般发票只能打印一次，但是因为打印机发票号码与待打印发票号码不一致，导致打印错误，需要重新打印
func (r *AlibabaeinvoicepaperprintAPIRequest) SetForcePrint(_forcePrint bool) error {
	r._forcePrint = _forcePrint
	r.Set("force_print", _forcePrint)
	return nil
}

// GetForcePrint ForcePrint Getter
func (r AlibabaeinvoicepaperprintAPIRequest) GetForcePrint() bool {
	return r._forcePrint
}
