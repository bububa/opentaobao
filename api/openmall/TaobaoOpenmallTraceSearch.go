package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltracesearch 获取Openmall订单物流流转信息
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
func Taobaoopenmalltracesearch(clt *core.SDKClient, req *openmall.TaobaoopenmalltracesearchAPIRequest, session string) (*openmall.TaobaoopenmalltracesearchAPIResponse, error) {
	var resp openmall.TaobaoopenmalltracesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
