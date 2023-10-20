package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanTaokeTrace 百川淘客打点
// taobao.baichuan.taoke.trace
//
// 百川淘客打点
func TaobaoBaichuanTaokeTrace(clt *core.SDKClient, req *baichuan.TaobaoBaichuanTaokeTraceAPIRequest, resp *baichuan.TaobaoBaichuanTaokeTraceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
