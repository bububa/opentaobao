package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生成打印机渲染命令（通过打印机命令打印） APIResponse
cainiao.cloudprint.cmdprint.render

isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
*/
type CainiaoCloudprintCmdprintRenderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cloudprint_cmdprint_render_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0成功，非0失败
    
    RetCode   string `json:"ret_code,omitempty" xml:"