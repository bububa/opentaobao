package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalnewdraftstandpointqueryAPIRequest 未采纳口径查询(新) API请求
// alibaba.legal.newdraftstandpoint.query
//
// 未采纳口径查询(新)
type AlibabalegalnewdraftstandpointqueryAPIRequest struct {
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

// NewAlibabalegalnewdraftstandpointqueryRequest 初始化AlibabalegalnewdraftstandpointqueryAPIRequest对象
func NewAlibabalegalnewdraftstandpointqueryRequest() *AlibabalegalnewdraftstandpointqueryAPIRequest {
	return &AlibabalegalnewdraftstandpointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.newdraftstandpoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusiId is BusiId Setter
// 业务系统实体id
func (r *AlibabalegalnewdraftstandpointqueryAPIRequest) SetBusiId(_busiId string) error {
	r._busiId = _busiId
	r.Set("busi_id", _busiId)
	return nil
}

// GetBusiId BusiId Getter
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetBusiId() string {
	return r._busiId
}

// SetInputSystemCode is InputSystemCode Setter
// 系统标识
func (r *AlibabalegalnewdraftstandpointqueryAPIRequest) SetInputSystemCode(_inputSystemCode string) error {
	r._inputSystemCode = _inputSystemCode
	r.Set("input_system_code", _inputSystemCode)
	return nil
}

// GetInputSystemCode InputSystemCode Getter
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetInputSystemCode() string {
	return r._inputSystemCode
}

// SetPageNum is PageNum Setter
// 页号
func (r *AlibabalegalnewdraftstandpointqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabalegalnewdraftstandpointqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabalegalnewdraftstandpointqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
