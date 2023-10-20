package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantGatewayCallbackAPIRequest 人卡关系回调 API请求
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
type AlibabaCampusGuardantGatewayCallbackAPIRequest struct {
	model.Params
	// md5
	_token string
	// 请求数据
	_data string
}

// NewAlibabaCampusGuardantGatewayCallbackRequest 初始化AlibabaCampusGuardantGatewayCallbackAPIRequest对象
func NewAlibabaCampusGuardantGatewayCallbackRequest() *AlibabaCampusGuardantGatewayCallbackAPIRequest {
	return &AlibabaCampusGuardantGatewayCallbackAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusGuardantGatewayCallbackAPIRequest) Reset() {
	r._token = ""
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.gateway.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// md5
func (r *AlibabaCampusGuardantGatewayCallbackAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetToken() string {
	return r._token
}

// SetData is Data Setter
// 请求数据
func (r *AlibabaCampusGuardantGatewayCallbackAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetData() string {
	return r._data
}

var poolAlibabaCampusGuardantGatewayCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusGuardantGatewayCallbackRequest()
	},
}

// GetAlibabaCampusGuardantGatewayCallbackRequest 从 sync.Pool 获取 AlibabaCampusGuardantGatewayCallbackAPIRequest
func GetAlibabaCampusGuardantGatewayCallbackAPIRequest() *AlibabaCampusGuardantGatewayCallbackAPIRequest {
	return poolAlibabaCampusGuardantGatewayCallbackAPIRequest.Get().(*AlibabaCampusGuardantGatewayCallbackAPIRequest)
}

// ReleaseAlibabaCampusGuardantGatewayCallbackAPIRequest 将 AlibabaCampusGuardantGatewayCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusGuardantGatewayCallbackAPIRequest(v *AlibabaCampusGuardantGatewayCallbackAPIRequest) {
	v.Reset()
	poolAlibabaCampusGuardantGatewayCallbackAPIRequest.Put(v)
}
