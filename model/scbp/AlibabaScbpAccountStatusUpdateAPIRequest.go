package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountStatusUpdateAPIRequest 修改账户级别关键词推广状态 API请求
// alibaba.scbp.account.status.update
//
// 修改账户级别关键词推广状态
type AlibabaScbpAccountStatusUpdateAPIRequest struct {
	model.Params
	// on:开启,off:暂停
	_status string
}

// NewAlibabaScbpAccountStatusUpdateRequest 初始化AlibabaScbpAccountStatusUpdateAPIRequest对象
func NewAlibabaScbpAccountStatusUpdateRequest() *AlibabaScbpAccountStatusUpdateAPIRequest {
	return &AlibabaScbpAccountStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountStatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStatus is Status Setter
// on:开启,off:暂停
func (r *AlibabaScbpAccountStatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaScbpAccountStatusUpdateAPIRequest) GetStatus() string {
	return r._status
}
