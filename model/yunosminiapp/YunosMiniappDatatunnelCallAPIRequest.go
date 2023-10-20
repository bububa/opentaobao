package yunosminiapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosMiniappDatatunnelCallAPIRequest 车载小程序外部服务调用 API请求
// yunos.miniapp.datatunnel.call
//
// 对客户提供的api进行统一封装调用。
type YunosMiniappDatatunnelCallAPIRequest struct {
	model.Params
	// 参数
	_param *YunosMiniappDatatunnelCallBaseRequest
}

// NewYunosMiniappDatatunnelCallRequest 初始化YunosMiniappDatatunnelCallAPIRequest对象
func NewYunosMiniappDatatunnelCallRequest() *YunosMiniappDatatunnelCallAPIRequest {
	return &YunosMiniappDatatunnelCallAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosMiniappDatatunnelCallAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosMiniappDatatunnelCallAPIRequest) GetApiMethodName() string {
	return "yunos.miniapp.datatunnel.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosMiniappDatatunnelCallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosMiniappDatatunnelCallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *YunosMiniappDatatunnelCallAPIRequest) SetParam(_param *YunosMiniappDatatunnelCallBaseRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r YunosMiniappDatatunnelCallAPIRequest) GetParam() *YunosMiniappDatatunnelCallBaseRequest {
	return r._param
}

var poolYunosMiniappDatatunnelCallAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosMiniappDatatunnelCallRequest()
	},
}

// GetYunosMiniappDatatunnelCallRequest 从 sync.Pool 获取 YunosMiniappDatatunnelCallAPIRequest
func GetYunosMiniappDatatunnelCallAPIRequest() *YunosMiniappDatatunnelCallAPIRequest {
	return poolYunosMiniappDatatunnelCallAPIRequest.Get().(*YunosMiniappDatatunnelCallAPIRequest)
}

// ReleaseYunosMiniappDatatunnelCallAPIRequest 将 YunosMiniappDatatunnelCallAPIRequest 放入 sync.Pool
func ReleaseYunosMiniappDatatunnelCallAPIRequest(v *YunosMiniappDatatunnelCallAPIRequest) {
	v.Reset()
	poolYunosMiniappDatatunnelCallAPIRequest.Put(v)
}
