package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkerquerypageAPIRequest 查询工人列表 API请求
// tmall.servicecenter.worker.querypage
//
// 服务商查询工人列表
type TmallservicecenterworkerquerypageAPIRequest struct {
	model.Params
	// 页码
	_pageIndex int64
}

// NewTmallservicecenterworkerquerypageRequest 初始化TmallservicecenterworkerquerypageAPIRequest对象
func NewTmallservicecenterworkerquerypageRequest() *TmallservicecenterworkerquerypageAPIRequest {
	return &TmallservicecenterworkerquerypageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkerquerypageAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.querypage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkerquerypageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkerquerypageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageIndex is PageIndex Setter
// 页码
func (r *TmallservicecenterworkerquerypageAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallservicecenterworkerquerypageAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
