package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
生成打印机渲染命令（通过打印机命令打印） APIResponse
cainiao.cloudprint.cmdprint.render

isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。
*/
type CainiaoCloudprintCmdprintRenderAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintCmdprintRenderResponse `json:"cainiao_cloudprint_cmdprint_render_response,omitempty"`
}

type CainiaoCloudprintCmdprintRenderResponse struct {

    // 0成功，非0失败
    RetCode   string `json:"ret_code,omitempty"`

    // 错误信息
    RetMsg   string `json:"ret_msg,omitempty"`

    // 指令集内容串
    CmdContent   string `json:"cmd_content,omitempty"`

    // 指令集编码方式：origin-原串 gzip-采用gzip压缩并base64编码
    CmdEncoding   string `json:"cmd_encoding,omitempty"`

}
