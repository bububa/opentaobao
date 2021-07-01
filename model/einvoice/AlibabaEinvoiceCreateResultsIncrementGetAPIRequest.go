package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceCreateResultsIncrementGetAPIRequest
ERP增量开票结果获取 API请求
alibaba.einvoice.create.results.increment.get

增量开票结果获取 */
type AlibabaEinvoiceCreateResultsIncrementGetAPIRequest struct {
	model.Params
	// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
	_status string
	// 起始查询时间
	_startModified string
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 终止查询时间
	_endModified string
	// 页面大小(不能超过200)
	_pageSize int64
	// 显示的页码
	_pageNo int64
}

// New
