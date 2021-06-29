package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生成打印机渲染命令（通过打印机命令打印） API请求
cainiao.cloudprint.cmdprint.render

isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
*/
type CainiaoCloudprintCmdprintRenderRequest struct {
    model.Params
    // 参数对象
    _params   *CmdRenderParams
}

// 初始化CainiaoCloudprintCmdprintRenderRequest对象
func NewCainiaoCloudprintCmdprintRenderRequest() *CainiaoCloudprintCmdprintRenderRequest{
    return &CainiaoCloudprintCmdprintRenderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCmdprintRenderRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.cmdprint.render"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCmdprintRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Params Setter
// 参数对象
func (r *CainiaoCloudprintCmdprintRenderRequest) SetParams(_params *CmdRenderParams) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r CainiaoCloudprintCmdprintRenderRequest) GetParams() *CmdRenderParams {
    return r._params
}
