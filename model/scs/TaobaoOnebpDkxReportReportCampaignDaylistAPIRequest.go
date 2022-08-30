package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest 获取计划分日报表 API请求
// taobao.onebp.dkx.report.report.campaign.daylist
//
// 获取计划分日报表，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe
type TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportCampaignDaylistRequest 初始化TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest对象
func NewTaobaoOnebpDkxReportReportCampaignDaylistRequest() *TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest {
	return &TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.campaign.daylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportCampaignDaylistAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}
