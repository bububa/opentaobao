package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointQuerystandpointAPIRequest 主动查询口径 API请求
// alibaba.legal.case.standpoint.querystandpoint
//
// 为法宝侧提供主动查询口径功能,有利于规范外部律师答辩口径.
type AlibabaLegalCaseStandpointQuerystandpointAPIRequest struct {
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

// NewAlibabaLegalCaseStandpointQuerystandpointRequest 初始化AlibabaLegalCaseStandpointQuerystandpointAPIRequest对象
func NewAlibabaLegalCaseStandpointQuerystandpointRequest() *AlibabaLegalCaseStandpointQuerystandpointAPIRequest {
	return &AlibabaLegalCaseStandpointQuerystandpointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.standpoint.querystandpoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 口径搜索关键字
func (r *AlibabaLegalCaseStandpointQuerystandpointAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetSupplierId is SupplierId Setter
// 供应商id
func (r *AlibabaLegalCaseStandpointQuerystandpointAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetSupplierId() string {
	return r._supplierId
}

// SetPageNum is PageNum Setter
// 当前页
func (r *AlibabaLegalCaseStandpointQuerystandpointAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaLegalCaseStandpointQuerystandpointAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaLegalCaseStandpointQuerystandpointAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
