package damaiticklet

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damaiticklet"
)

// Alibabadamaimxopengatewayscript 第三方剧本数据推送
// alibaba.damai.mx.opengateway.script
//
// 第三方剧本数据推送
func Alibabadamaimxopengatewayscript(clt *core.SDKClient, req *damaiticklet.AlibabadamaimxopengatewayscriptAPIRequest, session string) (*damaiticklet.AlibabadamaimxopengatewayscriptAPIResponse, error) {
	var resp damaiticklet.AlibabadamaimxopengatewayscriptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
