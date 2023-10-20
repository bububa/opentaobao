package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// CainiaoWaybillCloudprintNetprintPrint 网络打印机打印接口
// cainiao.waybill.cloudprint.netprint.print
//
// 打印接口
func CainiaoWaybillCloudprintNetprintPrint(clt *core.SDKClient, req *wlb.CainiaoWaybillCloudprintNetprintPrintAPIRequest, resp *wlb.CainiaoWaybillCloudprintNetprintPrintAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
