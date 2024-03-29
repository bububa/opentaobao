package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperPrintAPIRequest 纸票打印接口 API请求
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintAPIRequest struct {
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

// NewAlibabaEinvoicePaperPrintRequest 初始化AlibabaEinvoicePaperPrintAPIRequest对象
func NewAlibabaEinvoicePaperPrintRequest() *AlibabaEinvoicePaperPrintAPIRequest {
	return &AlibabaEinvoicePaperPrintAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoicePaperPrintAPIRequest) Reset() {
	r._payeeRegisterNo = ""
	r._serialNo = ""
	r._dialogSettingFlag = 0
	r._printFlag = 0
	r._forcePrint = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperPrintAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.print"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperPrintAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoicePaperPrintAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaEinvoicePaperPrintAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoicePaperPrintRequest()
	},
}

// GetAlibabaEinvoicePaperPrintRequest 从 sync.Pool 获取 AlibabaEinvoicePaperPrintAPIRequest
func GetAlibabaEinvoicePaperPrintAPIRequest() *AlibabaEinvoicePaperPrintAPIRequest {
	return poolAlibabaEinvoicePaperPrintAPIRequest.Get().(*AlibabaEinvoicePaperPrintAPIRequest)
}

// ReleaseAlibabaEinvoicePaperPrintAPIRequest 将 AlibabaEinvoicePaperPrintAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoicePaperPrintAPIRequest(v *AlibabaEinvoicePaperPrintAPIRequest) {
	v.Reset()
	poolAlibabaEinvoicePaperPrintAPIRequest.Put(v)
}
