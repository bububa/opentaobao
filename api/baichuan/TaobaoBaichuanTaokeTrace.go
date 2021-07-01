package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

/* TaobaoBaichuanTaokeTrace
百川淘客打点
taobao.baichuan.taoke.trace

百川淘客打点 */
func TaobaoBaichuanTaokeTrace(clt *core.SDKClient, req *baichuan.TaobaoBaichuanTaokeTraceAPIRequest, session string) (*baichuan.TaobaoBaichuanTaokeTraceAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanTaokeTraceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
