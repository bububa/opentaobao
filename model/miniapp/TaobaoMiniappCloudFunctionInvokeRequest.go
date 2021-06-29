package miniapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部触发云函数 API请求
taobao.miniapp.cloud.function.invoke

用户isv从外部触发聚石塔云函数的执行。
*/
type TaobaoMiniappCloudFunctionInvokeRequest struct {
    model.Params
    // 云函数名称
    _name   string
    // 指定云函数的handler
    _handler   string
    // 调用云函数时的传参（JSON格式）
    _data   string
    // 云环境
    _env   string
    // 扩展协议参数
    _exts   string
}

// 初始化TaobaoMiniappCloudFunctionInvokeRequest对象
func NewTaobaoMiniappCloudFunctionInvokeRequest() *TaobaoMiniappCloudFunctionInvokeRequest{
    return &TaobaoMiniappCloudFunctionInvokeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.function.invoke"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 云函数名称
func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetName() string {
    return r._name
}
// Handler Setter
// 指定云函数的handler
func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetHandler(_handler string) error {
    r._handler = _handler
    r.Set("handler", _handler)
    return nil
}

// Handler Getter
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetHandler() string {
    return r._handler
}
// Data Setter
// 调用云函数时的传参（JSON格式）
func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetData() string {
    return r._data
}
// Env Setter
// 云环境
func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetEnv(_env string) error {
    r._env = _env
    r.Set("env", _env)
    return nil
}

// Env Getter
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetEnv() string {
    return r._env
}
// Exts Setter
// 扩展协议参数
func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetExts(_exts string) error {
    r._exts = _exts
    r.Set("exts", _exts)
    return nil
}

// Exts Getter
func (r TaobaoMiniappCloudFunctionInvokeRequest) GetExts() string {
    return r._exts
}
