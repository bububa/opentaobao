package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFmhealthPressureReportCreateAPIRequest
血压报告接口 API请求
alibaba.fmhealth.pressure.report.create

生成用户血压测量报告 */
type AlibabaFmhealthPressureReportCreateAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
	// 报告类型
	_reportType string
	// 报告内容
	_reportData string
	// 报告周期
	_reportPeriod string
	// 报告时间
	_reportTime string
	// 报告周期天数
	_reportPeriodDays string
	// 数据来源
	_reportSource string
}

// New
