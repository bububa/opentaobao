package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasestandpointquerystandpointAPIRequest 主动查询口径 API请求
// alibaba.legal.case.standpoint.querystandpoint
//
// 为法宝侧提供主动查询口径功能,有利于规范外部律师答辩口径.
type AlibabalegalcasestandpointquerystandpointAPIRequest struct {
	model.Params
	// 口径搜索关键字
	_keyword string
	// 供应商id
	_supplierId string
	// 当前页
	_pageNum int64
	// 页大小
	_pageSize int64
}

// NewAlibabalegalcasestandpointquerystandpointRequest 初始化AlibabalegalcasestandpointquerystandpointAPIRequest对象
func NewAlibabalegalcasestandpointquerystandpointRequest() *AlibabalegalcasestandpointquerystandpointAPIRequest {
	return &AlibabalegalcasestandpointquerystandpointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.standpoint.querystandpoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 口径搜索关键字
func (r *AlibabalegalcasestandpointquerystandpointAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetSupplierId is SupplierId Setter
// 供应商id
func (r *AlibabalegalcasestandpointquerystandpointAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetSupplierId() string {
	return r._supplierId
}

// SetPageNum is PageNum Setter
// 当前页
func (r *AlibabalegalcasestandpointquerystandpointAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabalegalcasestandpointquerystandpointAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabalegalcasestandpointquerystandpointAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
