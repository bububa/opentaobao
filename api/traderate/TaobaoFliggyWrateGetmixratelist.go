package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// TaobaoFliggyWrateGetmixratelist 飞猪通用评价接口
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
func TaobaoFliggyWrateGetmixratelist(clt *core.SDKClient, req *traderate.TaobaoFliggyWrateGetmixratelistAPIRequest, session string) (*traderate.TaobaoFliggyWrateGetmixratelistAPIResponse, error) {
	var resp traderate.TaobaoFliggyWrateGetmixratelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
