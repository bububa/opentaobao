package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterContractsSearchAPIRequest
获取合同类的服务工单信息 API请求
tmall.servicecenter.contracts.search

获取合同类的服务工单信息 */
type TmallServicecenterContractsSearchAPIRequest struct {
	model.Params
	// 开始时间:  开始时间和结束时间不能超过15分钟
	_start int64
	// 结束时间:  开始时间和结束时间不能超过15分钟
	_end int64
}

// NewTmallServicecenterContractsSearchRequest 初始化TmallServicecenterContractsSearchAPIRequest对象
func NewTmallServicecenterContractsSearchRequest() *TmallServicecenterContractsSearchAPIRequest {
	return &TmallServicecenterContractsSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterContractsSearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.contracts.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterContractsSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// Get Start Getter
func (r TmallServicecenterContractsSearchAPIRequest) GetStart() int64 {
	return r._start
}

// Set is End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// Get End Getter
func (r TmallServicecenterContractsSearchAPIRequest) GetEnd() int64 {
	return r._end
}
