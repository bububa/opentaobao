package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePaperPrintAPIRequest
纸票打印接口 API请求
alibaba.einvoice.paper.print

打印一张已开具成功的纸票 */
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

// New
