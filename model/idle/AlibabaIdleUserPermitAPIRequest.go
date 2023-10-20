package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitAPIRequest 用户appkey授权 API请求
// alibaba.idle.user.permit
//
// 闲鱼卖家与服务商关系绑定，用于业务数据分发/授权校验/消息通知鉴权
type AlibabaIdleUserPermitAPIRequest struct {
	model.Params
	// 授权请求
	_paramUserGrantRequest *UserGrantRequest
}

// NewAlibabaIdleUserPermitRequest 初始化AlibabaIdleUserPermitAPIRequest对象
func NewAlibabaIdleUserPermitRequest() *AlibabaIdleUserPermitAPIRequest {
	return &AlibabaIdleUserPermitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleUserPermitAPIRequest) Reset() {
	r._paramUserGrantRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleUserPermitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUserGrantRequest is ParamUserGrantRequest Setter
// 授权请求
func (r *AlibabaIdleUserPermitAPIRequest) SetParamUserGrantRequest(_paramUserGrantRequest *UserGrantRequest) error {
	r._paramUserGrantRequest = _paramUserGrantRequest
	r.Set("param_user_grant_request", _paramUserGrantRequest)
	return nil
}

// GetParamUserGrantRequest ParamUserGrantRequest Getter
func (r AlibabaIdleUserPermitAPIRequest) GetParamUserGrantRequest() *UserGrantRequest {
	return r._paramUserGrantRequest
}

var poolAlibabaIdleUserPermitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleUserPermitRequest()
	},
}

// GetAlibabaIdleUserPermitRequest 从 sync.Pool 获取 AlibabaIdleUserPermitAPIRequest
func GetAlibabaIdleUserPermitAPIRequest() *AlibabaIdleUserPermitAPIRequest {
	return poolAlibabaIdleUserPermitAPIRequest.Get().(*AlibabaIdleUserPermitAPIRequest)
}

// ReleaseAlibabaIdleUserPermitAPIRequest 将 AlibabaIdleUserPermitAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleUserPermitAPIRequest(v *AlibabaIdleUserPermitAPIRequest) {
	v.Reset()
	poolAlibabaIdleUserPermitAPIRequest.Put(v)
}
