package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerQuerypageAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.querypage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerQuerypageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
