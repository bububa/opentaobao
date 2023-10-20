package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportaccountdaylistAPIRequest 获取账户分日报表 API请求
// taobao.onebp.dkx.report.report.account.daylist
//
// 获取账户分日报表
type TaobaoonebpdkxreportreportaccountdaylistAPIRequest struct {
	model.Params
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoonebpdkxreportreportaccountdaylistRequest 初始化TaobaoonebpdkxreportreportaccountdaylistAPIRequest对象
func NewTaobaoonebpdkxreportreportaccountdaylistRequest() *TaobaoonebpdkxreportreportaccountdaylistAPIRequest {
	return &TaobaoonebpdkxreportreportaccountdaylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxreportreportaccountdaylistAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.account.daylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxreportreportaccountdaylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxreportreportaccountdaylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoonebpdkxreportreportaccountdaylistAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoonebpdkxreportreportaccountdaylistAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
