package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountstatusupdateAPIRequest 修改账户级别关键词推广状态 API请求
// alibaba.scbp.account.status.update
//
// 修改账户级别关键词推广状态
type AlibabascbpaccountstatusupdateAPIRequest struct {
	model.Params
	// on:开启,off:暂停
	_status string
}

// NewAlibabascbpaccountstatusupdateRequest 初始化AlibabascbpaccountstatusupdateAPIRequest对象
func NewAlibabascbpaccountstatusupdateRequest() *AlibabascbpaccountstatusupdateAPIRequest {
	return &AlibabascbpaccountstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpaccountstatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpaccountstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpaccountstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// on:开启,off:暂停
func (r *AlibabascbpaccountstatusupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabascbpaccountstatusupdateAPIRequest) GetStatus() string {
	return r._status
}
