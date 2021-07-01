package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveReportAPIRequest
体检机构对接_体检报告查询 API请求
alibaba.alihealth.examination.reserve.report

体检机构对接_体检报告获取 */
type AlibabaAlihealthExaminationReserveReportAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 到检唯一标识
	_checkNo string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
	// 查询报告卡号
	_searchNo string
	// 查询报告密码
	_searchPwd string
}

// New
