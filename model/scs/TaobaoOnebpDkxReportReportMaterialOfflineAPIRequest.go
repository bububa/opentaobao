package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest 查询某计划分商品离线报表 API请求
// taobao.onebp.dkx.report.report.material.offline
//
// 查询某计划分商品离线报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-23&#34;,&#34;campaign_id_list&#34;:[2853805001],&#34;end_time&#34;:&#34;2021-09-24&#34;,&#34;effect&#34;: 15   }
type TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportMaterialOfflineRequest 初始化TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest对象
func NewTaobaoOnebpDkxReportReportMaterialOfflineRequest() *TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest {
	return &TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._reportQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.material.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}

var poolTaobaoOnebpDkxReportReportMaterialOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxReportReportMaterialOfflineRequest()
	},
}

// GetTaobaoOnebpDkxReportReportMaterialOfflineRequest 从 sync.Pool 获取 TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest
func GetTaobaoOnebpDkxReportReportMaterialOfflineAPIRequest() *TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest {
	return poolTaobaoOnebpDkxReportReportMaterialOfflineAPIRequest.Get().(*TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest)
}

// ReleaseTaobaoOnebpDkxReportReportMaterialOfflineAPIRequest 将 TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportMaterialOfflineAPIRequest(v *TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportMaterialOfflineAPIRequest.Put(v)
}
