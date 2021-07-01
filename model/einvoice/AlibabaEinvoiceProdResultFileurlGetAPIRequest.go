package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceProdResultFileurlGetAPIRequest
发票中台-发票文件下载地址查询 API请求
alibaba.einvoice.prod.result.fileurl.get

发票文件下载地址查询，外部ISV通过该接口可以查对应发票文件 */
type AlibabaEinvoiceProdResultFileurlGetAPIRequest struct {
	model.Params
	// 业务平台商户ID/卖家用户ID
	_platformUserId string
	// 发票号码
	_invoiceNo string
	// 发票代码
	_invoiceCode string
	// 发票文件类型，小写，pdf/ofd/jpg
	_fileType string
	// 业务平台code, 由发票中台分配
	_platformCode string
}

// New
