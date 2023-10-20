package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotbusinessrecipegetpageAPIRequest 分页查询食谱 API请求
// alibaba.ailabs.iot.business.recipe.getpage
//
// 分页查询食谱数据
type AlibabaailabsiotbusinessrecipegetpageAPIRequest struct {
	model.Params
	// 开放账号id
	_openAccountId string
	// 分页页码
	_pageNum int64
	// 分页大小
	_pageSize int64
}

// NewAlibabaailabsiotbusinessrecipegetpageRequest 初始化AlibabaailabsiotbusinessrecipegetpageAPIRequest对象
func NewAlibabaailabsiotbusinessrecipegetpageRequest() *AlibabaailabsiotbusinessrecipegetpageAPIRequest {
	return &AlibabaailabsiotbusinessrecipegetpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotbusinessrecipegetpageAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.getpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotbusinessrecipegetpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotbusinessrecipegetpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAccountId is OpenAccountId Setter
// 开放账号id
func (r *AlibabaailabsiotbusinessrecipegetpageAPIRequest) SetOpenAccountId(_openAccountId string) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// GetOpenAccountId OpenAccountId Getter
func (r AlibabaailabsiotbusinessrecipegetpageAPIRequest) GetOpenAccountId() string {
	return r._openAccountId
}

// SetPageNum is PageNum Setter
// 分页页码
func (r *AlibabaailabsiotbusinessrecipegetpageAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaailabsiotbusinessrecipegetpageAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *AlibabaailabsiotbusinessrecipegetpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaailabsiotbusinessrecipegetpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
