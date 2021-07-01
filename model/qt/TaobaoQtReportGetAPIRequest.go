package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQtReportGetAPIRequest
查询质检报告 API请求
taobao.qt.report.get

质检报告查询 */
type TaobaoQtReportGetAPIRequest struct {
	model.Params
	// 质检编号
	_qtCode string
}

// New
