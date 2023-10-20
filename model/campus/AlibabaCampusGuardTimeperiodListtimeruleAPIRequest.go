package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardTimeperiodListtimeruleAPIRequest 门禁控制器查询时间规则 API请求
// alibaba.campus.guard.timeperiod.listtimerule
//
// 门禁控制器查询时间规则
type AlibabaCampusGuardTimeperiodListtimeruleAPIRequest struct {
	model.Params
	// 时间规则查询条件参数
	_timePeriodQuery *TimePeriodQuery
}

// NewAlibabaCampusGuardTimeperiodListtimeruleRequest 初始化AlibabaCampusGuardTimeperiodListtimeruleAPIRequest对象
func NewAlibabaCampusGuardTimeperiodListtimeruleRequest() *AlibabaCampusGuardTimeperiodListtimeruleAPIRequest {
	return &AlibabaCampusGuardTimeperiodListtimeruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardTimeperiodListtimeruleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.timeperiod.listtimerule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardTimeperiodListtimeruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusGuardTimeperiodListtimeruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTimePeriodQuery is TimePeriodQuery Setter
// 时间规则查询条件参数
func (r *AlibabaCampusGuardTimeperiodListtimeruleAPIRequest) SetTimePeriodQuery(_timePeriodQuery *TimePeriodQuery) error {
	r._timePeriodQuery = _timePeriodQuery
	r.Set("time_period_query", _timePeriodQuery)
	return nil
}

// GetTimePeriodQuery TimePeriodQuery Getter
func (r AlibabaCampusGuardTimeperiodListtimeruleAPIRequest) GetTimePeriodQuery() *TimePeriodQuery {
	return r._timePeriodQuery
}
