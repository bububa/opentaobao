package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitembrandqueryAPIRequest 品牌信息查询 API请求
// alibaba.wdk.item.brand.query
//
// 品牌信息查询
type AlibabawdkitembrandqueryAPIRequest struct {
	model.Params
	// 查询关键词，不填则查询全部
	_keyword string
	// 起始位置
	_offset int64
	// 一页大小
	_pageSize int64
}

// NewAlibabawdkitembrandqueryRequest 初始化AlibabawdkitembrandqueryAPIRequest对象
func NewAlibabawdkitembrandqueryRequest() *AlibabawdkitembrandqueryAPIRequest {
	return &AlibabawdkitembrandqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitembrandqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.brand.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitembrandqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitembrandqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 查询关键词，不填则查询全部
func (r *AlibabawdkitembrandqueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabawdkitembrandqueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetOffset is Offset Setter
// 起始位置
func (r *AlibabawdkitembrandqueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r AlibabawdkitembrandqueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 一页大小
func (r *AlibabawdkitembrandqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabawdkitembrandqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
