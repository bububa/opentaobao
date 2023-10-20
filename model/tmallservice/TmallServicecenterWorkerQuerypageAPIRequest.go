package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerQuerypageAPIRequest 查询工人列表 API请求
// tmall.servicecenter.worker.querypage
//
// 服务商查询工人列表
type TmallServicecenterWorkerQuerypageAPIRequest struct {
	model.Params
	// 页码
	_pageIndex int64
}

// NewTmallServicecenterWorkerQuerypageRequest 初始化TmallServicecenterWorkerQuerypageAPIRequest对象
func NewTmallServicecenterWorkerQuerypageRequest() *TmallServicecenterWorkerQuerypageAPIRequest {
	return &TmallServicecenterWorkerQuerypageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkerQuerypageAPIRequest) Reset() {
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQuerypageAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.querypage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQuerypageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkerQuerypageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageIndex is PageIndex Setter
// 页码
func (r *TmallServicecenterWorkerQuerypageAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallServicecenterWorkerQuerypageAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

var poolTmallServicecenterWorkerQuerypageAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkerQuerypageRequest()
	},
}

// GetTmallServicecenterWorkerQuerypageRequest 从 sync.Pool 获取 TmallServicecenterWorkerQuerypageAPIRequest
func GetTmallServicecenterWorkerQuerypageAPIRequest() *TmallServicecenterWorkerQuerypageAPIRequest {
	return poolTmallServicecenterWorkerQuerypageAPIRequest.Get().(*TmallServicecenterWorkerQuerypageAPIRequest)
}

// ReleaseTmallServicecenterWorkerQuerypageAPIRequest 将 TmallServicecenterWorkerQuerypageAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkerQuerypageAPIRequest(v *TmallServicecenterWorkerQuerypageAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkerQuerypageAPIRequest.Put(v)
}
