package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportcampaigndaylistAPIRequest 获取计划分日报表 API请求
// taobao.onebp.dkx.report.report.campaign.daylist
//
// 获取计划分日报表，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe
type TaobaoonebpdkxreportreportcampaigndaylistAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoonebpdkxreportreportcampaigndaylistRequest 初始化TaobaoonebpdkxreportreportcampaigndaylistAPIRequest对象
func NewTaobaoonebpdkxreportreportcampaigndaylistRequest() *TaobaoonebpdkxreportreportcampaigndaylistAPIRequest {
	return &TaobaoonebpdkxreportreportcampaigndaylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.campaign.daylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoonebpdkxreportreportcampaigndaylistAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
