package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest 获取创意离线报表 API请求
// taobao.onebp.dkx.creative.creative.report.offline
//
// 获取创意离线报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 创意绑定的查询信息
	_creativeBindQuery *CreativeBindQueryTopDto
}

// NewTaobaoOnebpDkxCreativeCreativeReportOfflineRequest 初始化TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest对象
func NewTaobaoOnebpDkxCreativeCreativeReportOfflineRequest() *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest {
	return &TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.creative.creative.report.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetCreativeBindQuery is CreativeBindQuery Setter
// 创意绑定的查询信息
func (r *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest) SetCreativeBindQuery(_creativeBindQuery *CreativeBindQueryTopDto) error {
	r._creativeBindQuery = _creativeBindQuery
	r.Set("creative_bind_query", _creativeBindQuery)
	return nil
}

// GetCreativeBindQuery CreativeBindQuery Getter
func (r TaobaoOnebpDkxCreativeCreativeReportOfflineAPIRequest) GetCreativeBindQuery() *CreativeBindQueryTopDto {
	return r._creativeBindQuery
}
