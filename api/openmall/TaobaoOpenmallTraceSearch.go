package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTraceSearch 获取Openmall订单物流流转信息
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
func TaobaoOpenmallTraceSearch(clt *core.SDKClient, req *openmall.TaobaoOpenmallTraceSearchAPIRequest, session string) (*openmall.TaobaoOpenmallTraceSearchAPIResponse, error) {
	var resp openmall.TaobaoOpenmallTraceSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
