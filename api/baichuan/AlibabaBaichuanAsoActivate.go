package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanAsoActivate 设备安装活动激活
// alibaba.baichuan.aso.activate
//
// 设备安装活动激活
func AlibabaBaichuanAsoActivate(clt *core.SDKClient, req *baichuan.AlibabaBaichuanAsoActivateAPIRequest, resp *baichuan.AlibabaBaichuanAsoActivateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
