package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimIoscertProductionSet 设置ios证书
// taobao.openim.ioscert.production.set
//
// 设置ios证书
func TaobaoOpenimIoscertProductionSet(clt *core.SDKClient, req *openim.TaobaoOpenimIoscertProductionSetAPIRequest, session string) (*openim.TaobaoOpenimIoscertProductionSetAPIResponse, error) {
	var resp openim.TaobaoOpenimIoscertProductionSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
