package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuantaoketrace 百川淘客打点
// taobao.baichuan.taoke.trace
//
// 百川淘客打点
func Taobaobaichuantaoketrace(clt *core.SDKClient, req *baichuan.TaobaobaichuantaoketraceAPIRequest, session string) (*baichuan.TaobaobaichuantaoketraceAPIResponse, error) {
	var resp baichuan.TaobaobaichuantaoketraceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
