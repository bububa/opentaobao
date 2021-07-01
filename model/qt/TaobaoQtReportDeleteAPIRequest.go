package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQtReportDeleteAPIRequest
质检报告删除接口 API请求
taobao.qt.report.delete

删除质检报告 */
type TaobaoQtReportDeleteAPIRequest struct {
	model.Params
	// 一个质检服务唯一标识质量检验单的编号
	_qtCode string
}

// New
