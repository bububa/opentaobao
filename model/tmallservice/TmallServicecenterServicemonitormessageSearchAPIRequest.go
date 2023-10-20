package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicemonitormessageSearchAPIRequest 根据时间段查询服务商的服务预警消息列表(15分钟内) API请求
// tmall.servicecenter.servicemonitormessage.search
//
// 根据时间段查询服务商的服务预警消息列表(15分钟内)
type TmallServicecenterServicemonitormessageSearchAPIRequest struct {
	model.Params
	// 开始时间long值
	_start int64
	// 结束时间long值，与start相差15分钟
	_end int64
}

// NewTmallServicecenterServicemonitormessageSearchRequest 初始化TmallServicecenterServicemonitormessageSearchAPIRequest对象
func NewTmallServicecenterServicemonitormessageSearchRequest() *TmallServicecenterServicemonitormessageSearchAPIRequest {
	return &TmallServicecenterServicemonitormessageSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicemonitormessageSearchAPIRequest) Reset() {
	r._start = 0
	r._end = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicemonitormessageSearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicemonitormessageSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicemonitormessageSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间long值
func (r *TmallServicecenterServicemonitormessageSearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallServicecenterServicemonitormessageSearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间long值，与start相差15分钟
func (r *TmallServicecenterServicemonitormessageSearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallServicecenterServicemonitormessageSearchAPIRequest) GetEnd() int64 {
	return r._end
}

var poolTmallServicecenterServicemonitormessageSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicemonitormessageSearchRequest()
	},
}

// GetTmallServicecenterServicemonitormessageSearchRequest 从 sync.Pool 获取 TmallServicecenterServicemonitormessageSearchAPIRequest
func GetTmallServicecenterServicemonitormessageSearchAPIRequest() *TmallServicecenterServicemonitormessageSearchAPIRequest {
	return poolTmallServicecenterServicemonitormessageSearchAPIRequest.Get().(*TmallServicecenterServicemonitormessageSearchAPIRequest)
}

// ReleaseTmallServicecenterServicemonitormessageSearchAPIRequest 将 TmallServicecenterServicemonitormessageSearchAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicemonitormessageSearchAPIRequest(v *TmallServicecenterServicemonitormessageSearchAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicemonitormessageSearchAPIRequest.Put(v)
}
