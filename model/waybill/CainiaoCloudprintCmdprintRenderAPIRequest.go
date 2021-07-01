package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintCmdprintRenderAPIRequest
生成打印机渲染命令（通过打印机命令打印） API请求
cainiao.cloudprint.cmdprint.render

isv 进行无线打印，需要将渲染数据传给打印机，通过生成打印机命令的方式能够最大限度的减少移动设备和打印机之间通信数据量。 */
type CainiaoCloudprintCmdprintRenderAPIRequest struct {
	model.Params
	// 参数对象
	_params *CmdRenderParams
}

// New
