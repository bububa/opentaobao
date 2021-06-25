package waybill

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/waybill"
)

/* 
生成打印机渲染命令（通过打印机命令打印） 
cainiao.cloudprint.cmdprint.render

isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
*/
func CainiaoCloudprintCmdprintRender(clt *core.SDKClient, req *waybill.CainiaoCloudprintCmdprintRenderRequest, session string) (*waybill.CainiaoCloudprintCmdprintRenderResponse, error) {
    var resp waybill.CainiaoCloudprintCmdprintRenderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
