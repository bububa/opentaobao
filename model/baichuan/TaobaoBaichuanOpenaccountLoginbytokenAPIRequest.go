package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountLoginbytokenAPIRequest
百川TOKEN 登录 API请求
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录 */
type TaobaoBaichuanOpenaccountLoginbytokenAPIRequest struct {
	model.Params
	// name
	_name string
}

// NewTaobaoBaichuanOpenaccountLoginbytokenRequest 初始化TaobaoBaichuanOpenaccountLoginbytokenAPIRequest对象
func NewTaobaoBaichuanOpenaccountLoginbytokenRequest() *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest {
	return &TaobaoBaichuanOpenaccountLoginbytokenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetApiMethodName() string {
	return "taobao.baichuan.openaccount.loginbytoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoBaichuanOpenaccountLoginbytokenAPIRequest) GetName() string {
	return r._name
}
