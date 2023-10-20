package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardtimeperiodlisttimeruleAPIRequest 门禁控制器查询时间规则 API请求
// alibaba.campus.guard.timeperiod.listtimerule
//
// 门禁控制器查询时间规则
type AlibabacampusguardtimeperiodlisttimeruleAPIRequest struct {
	model.Params
	// 时间规则查询条件参数
	_timePeriodQuery *TimePeriodQuery
}

// NewAlibabacampusguardtimeperiodlisttimeruleRequest 初始化AlibabacampusguardtimeperiodlisttimeruleAPIRequest对象
func NewAlibabacampusguardtimeperiodlisttimeruleRequest() *AlibabacampusguardtimeperiodlisttimeruleAPIRequest {
	return &AlibabacampusguardtimeperiodlisttimeruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguardtimeperiodlisttimeruleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.timeperiod.listtimerule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguardtimeperiodlisttimeruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguardtimeperiodlisttimeruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTimePeriodQuery is TimePeriodQuery Setter
// 时间规则查询条件参数
func (r *AlibabacampusguardtimeperiodlisttimeruleAPIRequest) SetTimePeriodQuery(_timePeriodQuery *TimePeriodQuery) error {
	r._timePeriodQuery = _timePeriodQuery
	r.Set("time_period_query", _timePeriodQuery)
	return nil
}

// GetTimePeriodQuery TimePeriodQuery Getter
func (r AlibabacampusguardtimeperiodlisttimeruleAPIRequest) GetTimePeriodQuery() *TimePeriodQuery {
	return r._timePeriodQuery
}
