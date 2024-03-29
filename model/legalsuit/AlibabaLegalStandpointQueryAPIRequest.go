package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointQueryAPIRequest 口径查询 API请求
// alibaba.legal.standpoint.query
//
// 口径查询
type AlibabaLegalStandpointQueryAPIRequest struct {
	model.Params
	// 关键字
	_keyword string
	// 用户id
	_userId string
	// 系统标识
	_inputSystemCode string
	// 页号
	_pageNum int64
	// 页面大小
	_pageSize int64
}

// NewAlibabaLegalStandpointQueryRequest 初始化AlibabaLegalStandpointQueryAPIRequest对象
func NewAlibabaLegalStandpointQueryRequest() *AlibabaLegalStandpointQueryAPIRequest {
	return &AlibabaLegalStandpointQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalStandpointQueryAPIRequest) Reset() {
	r._keyword = ""
	r._userId = ""
	r._inputSystemCode = ""
	r._pageNum = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalStandpointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalStandpointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalStandpointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键字
func (r *AlibabaLegalStandpointQueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabaLegalStandpointQueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaLegalStandpointQueryAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLegalStandpointQueryAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalStandpointQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalStandpointQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetPageNum is PageNum Setter
// 页号
func (r *AlibabaLegalStandpointQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaLegalStandpointQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *AlibabaLegalStandpointQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaLegalStandpointQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaLegalStandpointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalStandpointQueryRequest()
	},
}

// GetAlibabaLegalStandpointQueryRequest 从 sync.Pool 获取 AlibabaLegalStandpointQueryAPIRequest
func GetAlibabaLegalStandpointQueryAPIRequest() *AlibabaLegalStandpointQueryAPIRequest {
	return poolAlibabaLegalStandpointQueryAPIRequest.Get().(*AlibabaLegalStandpointQueryAPIRequest)
}

// ReleaseAlibabaLegalStandpointQueryAPIRequest 将 AlibabaLegalStandpointQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalStandpointQueryAPIRequest(v *AlibabaLegalStandpointQueryAPIRequest) {
	v.Reset()
	poolAlibabaLegalStandpointQueryAPIRequest.Put(v)
}
