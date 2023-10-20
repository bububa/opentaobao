package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointqueryAPIRequest 口径查询 API请求
// alibaba.legal.standpoint.query
//
// 口径查询
type AlibabalegalstandpointqueryAPIRequest struct {
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

// NewAlibabalegalstandpointqueryRequest 初始化AlibabalegalstandpointqueryAPIRequest对象
func NewAlibabalegalstandpointqueryRequest() *AlibabalegalstandpointqueryAPIRequest {
	return &AlibabalegalstandpointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalstandpointqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.standpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalstandpointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalstandpointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键字
func (r *AlibabalegalstandpointqueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabalegalstandpointqueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabalegalstandpointqueryAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalegalstandpointqueryAPIRequest) GetUserId() string {
	return r._userId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalstandpointqueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalstandpointqueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetPageNum is PageNum Setter
// 页号
func (r *AlibabalegalstandpointqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabalegalstandpointqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *AlibabalegalstandpointqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabalegalstandpointqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
