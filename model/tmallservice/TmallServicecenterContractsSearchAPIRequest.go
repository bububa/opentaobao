package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterContractsSearchAPIRequest 获取合同类的服务工单信息 API请求
// tmall.servicecenter.contracts.search
//
// 获取合同类的服务工单信息
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterContractsSearchAPIRequest) Reset() {
	r._start = 0
	r._end = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterContractsSearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.contracts.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterContractsSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterContractsSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallServicecenterContractsSearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterContractsSearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallServicecenterContractsSearchAPIRequest) GetEnd() int64 {
	return r._end
}

var poolTmallServicecenterContractsSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterContractsSearchRequest()
	},
}

// GetTmallServicecenterContractsSearchRequest 从 sync.Pool 获取 TmallServicecenterContractsSearchAPIRequest
func GetTmallServicecenterContractsSearchAPIRequest() *TmallServicecenterContractsSearchAPIRequest {
	return poolTmallServicecenterContractsSearchAPIRequest.Get().(*TmallServicecenterContractsSearchAPIRequest)
}

// ReleaseTmallServicecenterContractsSearchAPIRequest 将 TmallServicecenterContractsSearchAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterContractsSearchAPIRequest(v *TmallServicecenterContractsSearchAPIRequest) {
	v.Reset()
	poolTmallServicecenterContractsSearchAPIRequest.Put(v)
}
