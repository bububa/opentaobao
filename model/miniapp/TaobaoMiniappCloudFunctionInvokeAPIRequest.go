package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudFunctionInvokeAPIRequest 外部触发云函数 API请求
// taobao.miniapp.cloud.function.invoke
//
// 用户isv从外部触发聚石塔云函数的执行。
type TaobaoMiniappCloudFunctionInvokeAPIRequest struct {
	model.Params
	// 云函数名称
	_name string
	// 指定云函数的handler
	_handler string
	// 调用云函数时的传参（JSON格式）
	_data string
	// 云环境
	_env string
	// 扩展协议参数
	_exts string
}

// NewTaobaoMiniappCloudFunctionInvokeRequest 初始化TaobaoMiniappCloudFunctionInvokeAPIRequest对象
func NewTaobaoMiniappCloudFunctionInvokeRequest() *TaobaoMiniappCloudFunctionInvokeAPIRequest {
	return &TaobaoMiniappCloudFunctionInvokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.function.invoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Name Setter
// 云函数名称
func (r *TaobaoMiniappCloudFunctionInvokeAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetName() string {
	return r._name
}

// Set is Handler Setter
// 指定云函数的handler
func (r *TaobaoMiniappCloudFunctionInvokeAPIRequest) SetHandler(_handler string) error {
	r._handler = _handler
	r.Set("handler", _handler)
	return nil
}

// Get Handler Getter
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetHandler() string {
	return r._handler
}

// Set is Data Setter
// 调用云函数时的传参（JSON格式）
func (r *TaobaoMiniappCloudFunctionInvokeAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetData() string {
	return r._data
}

// Set is Env Setter
// 云环境
func (r *TaobaoMiniappCloudFunctionInvokeAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// Get Env Getter
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetEnv() string {
	return r._env
}

// Set is Exts Setter
// 扩展协议参数
func (r *TaobaoMiniappCloudFunctionInvokeAPIRequest) SetExts(_exts string) error {
	r._exts = _exts
	r.Set("exts", _exts)
	return nil
}

// Get Exts Getter
func (r TaobaoMiniappCloudFunctionInvokeAPIRequest) GetExts() string {
	return r._exts
}
