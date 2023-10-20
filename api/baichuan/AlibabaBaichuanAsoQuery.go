package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanAsoQuery 查询app在设备上的安装信息
// alibaba.baichuan.aso.query
//
// 查询app在设备上的安装信息
func AlibabaBaichuanAsoQuery(clt *core.SDKClient, req *baichuan.AlibabaBaichuanAsoQueryAPIRequest, resp *baichuan.AlibabaBaichuanAsoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
