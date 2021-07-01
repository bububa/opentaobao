package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQtReportsGetAPIRequest
批量查询质检报告 API请求
taobao.qt.reports.get

批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告 */
type TaobaoQtReportsGetAPIRequest struct {
	model.Params
	// 质检服务商名
	_spName string
	// 质检类型，目前只支持查询qt_type=11的类型
	_qtType int64
	// 收费项code
	_servcieItemCode string
	// 送检者昵称
	_nick string
	// 查询时间段的开始时间
	_startTime string
	// 查询时间段的结束时间
	_endTime string
}

// New
