package miniapp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部触发云函数 APIRequest
taobao.miniapp.cloud.function.invoke

用户isv从外部触发聚石塔云函数的执行。
*/
type TaobaoMiniappCloudFunctionInvokeRequest struct {
    model.Params

    // 云函数名称
    name   string 

    // 指定云函数的handler
    handler   string 

    // 调用云函数时的传参（JSON格式）
    data   string 

    // 云环境
    env   string 

    // 扩展协议参数
    exts   string 

}

func NewTaobaoMiniappCloudFunctionInvokeRequest() *TaobaoMiniappCloudFunctionInvokeRequest{
    return &TaobaoMiniappCloudFunctionInvokeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetApiMethodName() string {
    return "taobao.miniapp.cloud.function.invoke"
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetName() string {
    return r.name
}

func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetHandler(handler string) error {
    r.handler = handler
    r.Set("handler", handler)
    return nil
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetHandler() string {
    return r.handler
}

func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetData() string {
    return r.data
}

func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetEnv(env string) error {
    r.env = env
    r.Set("env", env)
    return nil
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetEnv() string {
    return r.env
}

func (r *TaobaoMiniappCloudFunctionInvokeRequest) SetExts(exts string) error {
    r.exts = exts
    r.Set("exts", exts)
    return nil
}

func (r TaobaoMiniappCloudFunctionInvokeRequest) GetExts() string {
    return r.exts
}

