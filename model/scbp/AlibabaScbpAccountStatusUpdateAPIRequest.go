package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAccountStatusUpdateAPIRequest) Reset() {
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAccountStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpAccountStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAccountStatusUpdateRequest()
	},
}

// GetAlibabaScbpAccountStatusUpdateRequest 从 sync.Pool 获取 AlibabaScbpAccountStatusUpdateAPIRequest
func GetAlibabaScbpAccountStatusUpdateAPIRequest() *AlibabaScbpAccountStatusUpdateAPIRequest {
	return poolAlibabaScbpAccountStatusUpdateAPIRequest.Get().(*AlibabaScbpAccountStatusUpdateAPIRequest)
}

// ReleaseAlibabaScbpAccountStatusUpdateAPIRequest 将 AlibabaScbpAccountStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAccountStatusUpdateAPIRequest(v *AlibabaScbpAccountStatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaScbpAccountStatusUpdateAPIRequest.Put(v)
}
