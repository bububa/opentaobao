package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountRegisterAPIRequest
百川账号注册 API请求
taobao.baichuan.openaccount.register

百川账号注册 */
type TaobaoBaichuanOpenaccountRegisterAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountRegisterRequest 初始化TaobaoBaichuanOpenaccountRegisterAPIRequest对象
func NewTaobaoBaichuanOpenaccountRegisterRequest() *TaobaoBaichuanOpenaccountRegisterAPIRequest {
	return &TaobaoBaichuanOpenaccountRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountRegisterAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoBaichuanOpenaccountRegisterAPIRequest) GetName() string {
	return r._name
}
