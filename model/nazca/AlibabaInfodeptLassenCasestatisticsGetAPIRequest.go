package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInfodeptLassenCasestatisticsGetAPIRequest
法庭提交和结案案件量接口 API请求
alibaba.infodept.lassen.casestatistics.get

功能描述：云嘉为浙江省高院制作数据大屏，需展示网上法庭相关数据，我方为省高院提供浙江省内法院收案和结案的案件量，开放数据接口，供其调取这两组数据。 */
type AlibabaInfodeptLassenCasestatisticsGetAPIRequest struct {
	model.Params
	// 地区代码
	_areaCode string
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
}

// New
