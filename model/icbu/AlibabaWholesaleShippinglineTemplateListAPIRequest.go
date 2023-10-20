package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawholesaleshippinglinetemplatelistAPIRequest 获取运费模板 API请求
// alibaba.wholesale.shippingline.template.list
//
// 查询运费模板信息
type AlibabawholesaleshippinglinetemplatelistAPIRequest struct {
	model.Params
	// 第几页从1开始
	_pageNum int64
	// 每页返回的数据个数
	_count int64
}

// NewAlibabawholesaleshippinglinetemplatelistRequest 初始化AlibabawholesaleshippinglinetemplatelistAPIRequest对象
func NewAlibabawholesaleshippinglinetemplatelistRequest() *AlibabawholesaleshippinglinetemplatelistAPIRequest {
	return &AlibabawholesaleshippinglinetemplatelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawholesaleshippinglinetemplatelistAPIRequest) GetApiMethodName() string {
	return "alibaba.wholesale.shippingline.template.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawholesaleshippinglinetemplatelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawholesaleshippinglinetemplatelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNum is PageNum Setter
// 第几页从1开始
func (r *AlibabawholesaleshippinglinetemplatelistAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabawholesaleshippinglinetemplatelistAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetCount is Count Setter
// 每页返回的数据个数
func (r *AlibabawholesaleshippinglinetemplatelistAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r AlibabawholesaleshippinglinetemplatelistAPIRequest) GetCount() int64 {
	return r._count
}
