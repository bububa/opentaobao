package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalNewdraftstandpointQueryAPIRequest 未采纳口径查询(新) API请求
// alibaba.legal.newdraftstandpoint.query
//
// 未采纳口径查询(新)
type AlibabaLegalNewdraftstandpointQueryAPIRequest struct {
	model.Params
	// 业务系统实体id
	_busiId string
	// 系统标识
	_inputSystemCode string
	// 页号
	_pageNum int64
	// 页大小
	_pageSize int64
}

// NewAlibabaLegalNewdraftstandpointQueryRequest 初始化AlibabaLegalNewdraftstandpointQueryAPIRequest对象
func NewAlibabaLegalNewdraftstandpointQueryRequest() *AlibabaLegalNewdraftstandpointQueryAPIRequest {
	return &AlibabaLegalNewdraftstandpointQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalNewdraftstandpointQueryAPIRequest) Reset() {
	r._busiId = ""
	r._inputSystemCode = ""
	r._pageNum = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.newdraftstandpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusiId is BusiId Setter
// 业务系统实体id
func (r *AlibabaLegalNewdraftstandpointQueryAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabaLegalNewdraftstandpointQueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetPageNum is PageNum Setter
// 页号
func (r *AlibabaLegalNewdraftstandpointQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaLegalNewdraftstandpointQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaLegalNewdraftstandpointQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaLegalNewdraftstandpointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalNewdraftstandpointQueryRequest()
	},
}

// GetAlibabaLegalNewdraftstandpointQueryRequest 从 sync.Pool 获取 AlibabaLegalNewdraftstandpointQueryAPIRequest
func GetAlibabaLegalNewdraftstandpointQueryAPIRequest() *AlibabaLegalNewdraftstandpointQueryAPIRequest {
	return poolAlibabaLegalNewdraftstandpointQueryAPIRequest.Get().(*AlibabaLegalNewdraftstandpointQueryAPIRequest)
}

// ReleaseAlibabaLegalNewdraftstandpointQueryAPIRequest 将 AlibabaLegalNewdraftstandpointQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalNewdraftstandpointQueryAPIRequest(v *AlibabaLegalNewdraftstandpointQueryAPIRequest) {
	v.Reset()
	poolAlibabaLegalNewdraftstandpointQueryAPIRequest.Put(v)
}
