package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscUserCenterInfoQueryAPIRequest 查询授权的用户信息 API请求
// alibaba.alsc.user.center.info.query
//
// 获取授权的饿了么用户信息
type AlibabaAlscUserCenterInfoQueryAPIRequest struct {
	model.Params
	// 请求模型
	_alscUserQueryOpenRequest *AlscUserQueryOpenRequest
}

// NewAlibabaAlscUserCenterInfoQueryRequest 初始化AlibabaAlscUserCenterInfoQueryAPIRequest对象
func NewAlibabaAlscUserCenterInfoQueryRequest() *AlibabaAlscUserCenterInfoQueryAPIRequest {
	return &AlibabaAlscUserCenterInfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscUserCenterInfoQueryAPIRequest) Reset() {
	r._alscUserQueryOpenRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscUserCenterInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.user.center.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscUserCenterInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscUserCenterInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlscUserQueryOpenRequest is AlscUserQueryOpenRequest Setter
// 请求模型
func (r *AlibabaAlscUserCenterInfoQueryAPIRequest) SetAlscUserQueryOpenRequest(_alscUserQueryOpenRequest *AlscUserQueryOpenRequest) error {
	r._alscUserQueryOpenRequest = _alscUserQueryOpenRequest
	r.Set("alsc_user_query_open_request", _alscUserQueryOpenRequest)
	return nil
}

// GetAlscUserQueryOpenRequest AlscUserQueryOpenRequest Getter
func (r AlibabaAlscUserCenterInfoQueryAPIRequest) GetAlscUserQueryOpenRequest() *AlscUserQueryOpenRequest {
	return r._alscUserQueryOpenRequest
}

var poolAlibabaAlscUserCenterInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscUserCenterInfoQueryRequest()
	},
}

// GetAlibabaAlscUserCenterInfoQueryRequest 从 sync.Pool 获取 AlibabaAlscUserCenterInfoQueryAPIRequest
func GetAlibabaAlscUserCenterInfoQueryAPIRequest() *AlibabaAlscUserCenterInfoQueryAPIRequest {
	return poolAlibabaAlscUserCenterInfoQueryAPIRequest.Get().(*AlibabaAlscUserCenterInfoQueryAPIRequest)
}

// ReleaseAlibabaAlscUserCenterInfoQueryAPIRequest 将 AlibabaAlscUserCenterInfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscUserCenterInfoQueryAPIRequest(v *AlibabaAlscUserCenterInfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlscUserCenterInfoQueryAPIRequest.Put(v)
}
